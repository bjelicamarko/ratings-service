package rabbitHandlers

// func unmarshal(message *models.Message, extractReference interface{}) error {
// 	jsonBytes, err := json.Marshal(message.Body)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(jsonBytes, &extractReference)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func publishMessage(repository *repositories.RatingsRepository, body interface{}, messageType models.MessageType, entityId uint) {
// 	bodyToPublish, err := json.Marshal(body)
// 	if err != nil {
// 		log.Printf("[CRITICAL] Could not marshall %s for entity with id: [%d]", messageType, entityId)
// 		return
// 	}

// 	message := models.Message{
// 		Type: messageType,
// 		Body: bodyToPublish,
// 	}

// 	err = repository.PublishMessage(&message)
// 	if err != nil {
// 		log.Printf("[CRITICAL] Could not publish %s for entity with id: [%d]! Error: %s", messageType, entityId, err.Error())
// 		return
// 	}
// }
