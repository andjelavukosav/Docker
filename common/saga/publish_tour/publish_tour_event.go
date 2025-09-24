package publish_tour

type TourDetails struct {
    ID          string
    Name        string
    AuthorID    string
    Description string
    Difficulty  string
    Price       float64
    Status      string
}

type PublishTourCommandType int8

const (
    PublishTour PublishTourCommandType = iota
    CancelTour
    UnknownCommand
)

type PublishTourCommand struct {
    Tour TourDetails
    Type PublishTourCommandType
}

type PublishTourReplyType int8

const (
    TourPublishedSuccessfully PublishTourReplyType = iota
    TourPublishFailed
    TourCancelled
    UnknownReply
)

type PublishTourReply struct {
    Tour TourDetails
    Type PublishTourReplyType
}
