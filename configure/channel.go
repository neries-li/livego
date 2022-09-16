package configure

type RoomKeysType struct {
}

var RoomKeys = &RoomKeysType{}

// set/reset a random key for channel
func (r *RoomKeysType) SetKey(channel string) (key string, err error) {
	return channel, nil
}

func (r *RoomKeysType) GetKey(channel string) (newKey string, err error) {
	return channel, nil
}

func (r *RoomKeysType) GetChannel(key string) (channel string, err error) {
	return key, nil
}

func (r *RoomKeysType) DeleteChannel(channel string) bool {
	return true
}

func (r *RoomKeysType) DeleteKey(key string) bool {
	return true
}
