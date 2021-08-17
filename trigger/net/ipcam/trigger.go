package ipcam

import (
	"context"
	"net/url"

	"github.com/r2d2-ai/aiflow/data/metadata"
	"github.com/r2d2-ai/aiflow/support/log"
	"github.com/r2d2-ai/aiflow/trigger"
	"gocv.io/x/gocv"
)

var triggerMd = trigger.NewMetadata(&Settings{}, &HandlerSettings{}, &Output{})

func init() {
	_ = trigger.Register(&Trigger{}, &Factory{})
}

type Factory struct {
}

// Metadata implements trigger.Factory.Metadata
func (*Factory) Metadata() *trigger.Metadata {
	return triggerMd
}

// New implements trigger.Factory.New
func (*Factory) New(config *trigger.Config) (trigger.Trigger, error) {
	s := &Settings{}
	err := metadata.MapToStruct(config.Settings, s, true)
	if err != nil {
		return nil, err
	}

	return &Trigger{settings: s}, nil
}

type Trigger struct {
	// metadata       *trigger.Metadata
	settings       *Settings
	cameraHandlers []*CameraHandler
	logger         log.Logger
}

type CameraHandler struct {
	shutdown chan bool
	handler  trigger.Handler
	settings *HandlerSettings
	cap      *gocv.VideoCapture
	logger   log.Logger
	id       string
}

func (t *Trigger) Initialize(ctx trigger.InitContext) error {

	t.logger = log.ChildLogger(ctx.Logger(), "ipcam-feed-listener")
	for _, handler := range ctx.GetHandlers() {
		s := &HandlerSettings{}
		err := metadata.MapToStruct(handler.Settings(), s, true)
		if err != nil {
			return err
		}

		if s.GroupId == "" {
			s.GroupId = s.Host
		}

		if s.CameraId == "" {
			s.CameraId = s.VideoURI
		}

		camHnd := &CameraHandler{}
		camHnd.settings = s
		camHnd.handler = handler
		camHnd.logger = t.logger
		t.cameraHandlers = append(t.cameraHandlers, camHnd)
	}
	return nil
}

// Start implements util.Managed.Start
func (t *Trigger) Start() error {

	for _, camHdl := range t.cameraHandlers {
		go camHdl.startStream()
	}
	return nil
}

func (camHnd *CameraHandler) startStream() error {
	//protocol := camHnd.settings.Protocol

	host := camHnd.settings.Host
	user := camHnd.settings.User
	password := camHnd.settings.Password
	videoUri := camHnd.settings.VideoURI
	id := camHnd.settings.GroupId + "/" + camHnd.settings.CameraId
	camHnd.logger.Infof("Start IP Cam %v stream", id)
	cap, err := gocv.OpenVideoCapture("rtsp://" + url.QueryEscape(user) + ":" + url.QueryEscape(password) + "@" + host + "/" + url.QueryEscape(videoUri))

	if err != nil {
		return err
	}

	camHnd.cap = cap
	camHnd.shutdown = make(chan bool)
	camHnd.id = id
	go camHnd.run()
	return nil
}

func (camHnd *CameraHandler) run() {
	var err error

	img := gocv.NewMat()
	defer img.Close()

	camHnd.logger.Infof("Running IP Cam %v stream", camHnd.id)
	//var times []int64
	for {
		//t_start := time.Now()
		select {
		case <-camHnd.shutdown:
			camHnd.logger.Infof("Stopping IP Cam %v stream", camHnd.id)
			camHnd.cap.Close()
			return
		default:
			camHnd.cap.Read(&img)
		}

		output := &Output{}
		output.Image = &img //send pointer
		output.CameraId = camHnd.settings.CameraId
		output.GroupdId = camHnd.settings.GroupId
		_, err = camHnd.handler.Handle(context.Background(), output)
		//t_duration := time.Since(t_start).Milliseconds()
		//times = append(times, t_duration)
		// if len(times) == 100 {
		// 	var sum int64 = 0
		// 	for _, t := range times {
		// 		sum += t
		// 	}
		// 	mean := float64(sum) / float64(len(times))
		// 	camHnd.logger.Infof("Average time %d", mean)
		// 	times = times[75:]
		// }
		if err != nil {
			camHnd.logger.Errorf("Failed to handle frame for IP Cam %v ", camHnd.id)
		}
	}
}

// Stop implements util.Managed.Stop
func (t *Trigger) Stop() error {
	for _, camHnd := range t.cameraHandlers {
		camHnd.shutdown <- true
	}
	return nil
}
