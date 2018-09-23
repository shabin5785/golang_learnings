package interfaces

//CarBuilderInterface is exported
type CarBuilderInterface interface {
	TopSpeed(int) CarBuilderInterface
	Paint(string) CarBuilderInterface
	Build() CarInterface
}
