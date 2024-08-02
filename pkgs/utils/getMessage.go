package utils

import (
	"github.com/golang/protobuf/proto"
	"strings"
	"wabot/models"
)

func GetMessage(reminder bool, client models.Client, service string, account models.Account) *string {

	trimService := strings.TrimSpace(service)
	trimName := strings.TrimSpace(client.Name)

	msg := proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nQueriamos recordate que *pronto vencerá la cuenta de " + trimService + "*, asociada al correo: " + account.Mail + ". \n\nPor favor, mantente al tanto de este aviso y asegurate de tener todo listo para el pago correspondiente. \n\n¡Gracias por tu continua confianza y preferencia! 🙌")

	if client.IsResellerCustomer {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nQueremos recordate que la cuenta de *" + trimService + "* a nombre de *" + trimName + "* está por vencer pronto.\n\nAgredecemos tu atención y getión al respecto.")
	}

	if reminder {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nSolo un recordatorio sobre su pago pendiente de *" + trimService + "*, asociado al correo: *" + account.Mail + "*. \n\nAgradecemos su pronta gestión. 🙌")
	}

	if reminder && client.IsResellerCustomer {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nSolo un recordatorio sobre el pago pendiente de *" + trimService + "* a nombre de *" + trimName + "*, asociado al correo: *" + account.Mail + "*. \n\nAgradecemos su pronta gestión. 🙌")
	}

	return msg
}
