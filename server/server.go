package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_stream-local/server/sensor"
	sensorpb "grpc_stream-local/server/sensorpb/protos"
	"log"
	"net"
	"time"
)

type server struct {
	Sensor *sensor.Sensor
}

func (s *server) TempSensor(req *sensorpb.SensorRequest,
	stream sensorpb.Sensor_TempSensorServer) error {
	for {
		temp := s.Sensor.GetTempSensor()
		err := stream.Send(&sensorpb.SensorResponse{Value: temp})
		if err != nil {
			log.Println("Error sending metric message ", err)
		}
		time.Sleep(time.Second * 5)
	}
	return nil
}

func (s *server) HumiditySensor(req *sensorpb.SensorRequest,
	stream sensorpb.Sensor_HumiditySensorServer) error {

	for {

		humd := s.Sensor.GetHumiditySensor()

		err := stream.Send(&sensorpb.SensorResponse{Value: humd})
		if err != nil {
			log.Println("Error sending metric message ", err)
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

func (s *server) ToxicitySensor(req *sensorpb.SensorRequest,
	stream sensorpb.Sensor_ToxicitySensorServer) error {
	for {
		toxic := s.Sensor.GetToxicity()
		err := stream.Send(&sensorpb.SensorResponse{Value: toxic})
		if err != nil {
			log.Println("Error sending metric message ", err)
		}
		time.Sleep(time.Second * 2)
	}
}

func (s *server) SetToxicityThreshold(ctx context.Context, res *sensorpb.ThresholdRequest) (*sensorpb.ThresholdResponse, error) {
	fmt.Println("SetToxicityThreshold called!!!", res)
	val := s.Sensor.SetToxicitySensor(int(res.Value))
	return &sensorpb.ThresholdResponse{
		Value: int64(val),
	}, nil
}

var (
	port int = 8080
)

func main() {

	sns := sensor.NewSensor()

	sns.StartMonitoring()

	addr := fmt.Sprintf(":%d", port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	sensorpb.RegisterSensorServer(s, &server{Sensor: sns})

	log.Printf("Starting server in port :%d\n", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}

}
