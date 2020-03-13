package kafka

// KafkaLibrary ...
type KafkaLibrary struct{}

// KafkaLibraryHandler ...
func KafkaLibraryHandler() *KafkaLibrary {
	return &KafkaLibrary{}
}

// KafkaLibraryInterface ...
type KafkaLibraryInterface interface {
	Init()
}

// Init ...
func Init() {

}
