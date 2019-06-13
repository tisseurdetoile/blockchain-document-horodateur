// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
	models	"github.com/geneva_horodateur/models"
//	models "github.com/Magicking/rc-ge-ch-pdf/models"
)

// MonitoringOKCode is the HTTP code returned for type MonitoringOK
const MonitoringOKCode int = 200

/*MonitoringOK Tout est en ordre et fonctionne correctement.


swagger:response monitoringOK
*/
type MonitoringOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Sonde `json:"body,omitempty"`
}

// NewMonitoringOK creates MonitoringOK with default headers values
func NewMonitoringOK() *MonitoringOK {

	return &MonitoringOK{}
}

// WithPayload adds the payload to the monitoring o k response
func (o *MonitoringOK) WithPayload(payload []*models.Sonde) *MonitoringOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the monitoring o k response
func (o *MonitoringOK) SetPayload(payload []*models.Sonde) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MonitoringOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Sonde, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*MonitoringDefault Internal error

swagger:response monitoringDefault
*/
type MonitoringDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewMonitoringDefault creates MonitoringDefault with default headers values
func NewMonitoringDefault(code int) *MonitoringDefault {
	if code <= 0 {
		code = 500
	}

	return &MonitoringDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the monitoring default response
func (o *MonitoringDefault) WithStatusCode(code int) *MonitoringDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the monitoring default response
func (o *MonitoringDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the monitoring default response
func (o *MonitoringDefault) WithPayload(payload *models.Error) *MonitoringDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the monitoring default response
func (o *MonitoringDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *MonitoringDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
