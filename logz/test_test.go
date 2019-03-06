package logz

func newFrame() *Frame {
	return &Frame{
		FrameId:           "bar",
		ParentFrameId: "baz",
		FrameName:         "bell",
		StackId:           "bang",
	}
}
