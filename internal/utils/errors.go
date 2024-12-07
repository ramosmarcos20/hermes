package utils

import "errors"

var (

	InvalidInput 	= errors.New("Input Invalido")
	ErrorDataBase	= errors.New("Error de Base de datos")

	/* USERS */
	UserNotFound 	= errors.New("Usuario no encontrado")
	EmailInUse		= errors.New("Este Email ya esta registrado")
	UserNameInUse	= errors.New("UserName ya esta registrado")

	/* TENANTS */
	TaxInUse		= errors.New("Este DNI ya esta registrado")
)