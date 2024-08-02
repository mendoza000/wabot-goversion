package utils

import (
	"github.com/golang/protobuf/proto"
	"wabot/models"
)

func GetMessage(reminder bool, client models.Client, service string, account models.Account) *string {

	msg := proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nQueriamos recordate que *pronto vencerá la cuenta de " + service + "*, asociada al correo: " + account.Mail + ". \n\nPor favor, mantente al tanto de este aviso y asegurate de tener todo listo para el pago correspondiente. \n\n¡Gracias por tu continua confianza y preferencia! 🙌")

	if client.IsResellerCustomer {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nQueremos recordate que la cuenta de *" + service + "* a nombre de *" + client.Name + "* está por vecer pronto.\n\nAgredecemos tu atención y getión al respecto.")
	}

	if reminder {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nSolo un recordatorio sobre su pago pendiente de *" + service + "*, asociado al correo: *" + account.Mail + "*. \n\nAgradecemos su pronta gestión. 🙌")
	}

	if reminder && client.IsResellerCustomer {
		msg = proto.String("*Mensaje automatizado* \n\nHola, buen dia 🙌. \nSolo un recordatorio sobre el pago pendiente de *" + service + "* a nombre de *" + client.Name + "*, asociado al correo: *" + account.Mail + "*. \n\nAgradecemos su pronta gestión. 🙌")
	}

	return msg
}
