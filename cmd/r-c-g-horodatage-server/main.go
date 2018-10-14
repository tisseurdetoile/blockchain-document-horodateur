// Code generated by go-swagger; DO NOT EDIT.

package main

import (
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"github.com/Magicking/rc-ge-ch-pdf/restapi"
	"github.com/Magicking/rc-ge-ch-pdf/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewRCGHorodatageAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "RCG horodatage"
	parser.LongDescription = "RCG horodatage est un service qui permet l'horodatage numérique via\nsur la blockchain Ethereum.\nLe principe est d'envoyer des fichiers qui sont ensuite passer dans\nune fonction hachage SHA3-256. Les « hash » sont ensuite intégrés\ndans un arbre de Merkle dont la racine est inséré dans une\ntransaction blockchain, l'(es) adresse(s) signant la transaction\nidentifie le Registre du Commerce, c'est une information qui doit\nêtre publique.\n"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
