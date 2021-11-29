package main

import (
	"awesomeProject2/api"
	"awesomeProject2/models"
	"awesomeProject2/utils"
	"fmt"
)

func MenuHome(comm *api.Communication) {
	for {
		fmt.Println("\nBenvenuto nel simulatore di investimenti all'Edge!")
		fmt.Println("Scegli un opzione:")
		var v int
		//we need a do-while behaviour, the following is a way to realize it
		for {
			fmt.Println("1)Registrati")
			fmt.Println("2)Login")
			fmt.Scan(&v)
			if v >= 1 || v <= 2 {
				break
			} else {
				fmt.Println("Scelta ERRATA! Ripeti la scelta..")
			}
		}

		switch v {
		case 1:
			fmt.Println("Inserisci Nome")
			var firstName string
			fmt.Scan(&firstName)
			fmt.Println("Inserisci Cognome")
			var name string
			fmt.Scan(&name)
			fmt.Println("Inserisci username")
			var username string
			fmt.Scan(&username)
			fmt.Println("Inserisci password")
			var password string
			fmt.Scan(&password)
			userCreateReq := models.UserReq{FirstName: firstName, Name: name, Username: username, Password: password}
			res := comm.CreateUser(userCreateReq)
			if res["detail"] == "Username already registered" {
				fmt.Println("Utente già registrato!")
			} else {
				utils.PrintMap(res)
			}

		case 2:
			fmt.Println("Inserisci username")
			var username string
			fmt.Scan(&username)
			fmt.Println("Inserisci password")
			var password string
			fmt.Scan(&password)
			res := comm.Login(models.LoginReq{Username: username, Password: password})
			if res != "" {
				fmt.Println("Login effettuato con successo!")

				//we register in communication the current user logged and the token
				comm.SetToken(res)
				currentUser := comm.GetReqUser(username)
				comm.SetCurrentUser(currentUser)
				//to return the control to main
				Menu1(comm)
			} else {
				fmt.Println("Login non effettuato!")
			}
		}
	}
}

func Menu1(comm *api.Communication) {
	for {
		var v int
		for {
			fmt.Println("\n1)Elimina Utente")
			fmt.Println("2)Prosegui con la simulazione")
			fmt.Println("3)Indietro")
			fmt.Scan(&v)
			if v >= 1 || v <= 2 {
				break
			} else {
				fmt.Println("Scelta ERRATA! Ripeti la scelta..")
			}
		}

		switch v {
		case 1:
			if comm.GetToken() == ""{
				fmt.Println("Devi effettuare il Login!")
				break
			}
			fmt.Println("Inserisci username")
			var username string
			fmt.Scan(&username)
			fmt.Println("Inserisci password")
			var password string
			fmt.Scan(&password)
			tmp:=comm.DeleteUser(models.LoginReq{Username: username, Password: password})

			switch tmp {
			case "User not found!":
				fmt.Println("Username or Password non presenti!")
			case "Token expired redo the login!":
				fmt.Println("Il token è scaduto: rifai il login")
			case "Incorrect username or password":
				fmt.Println("Username or Password non errati")
			case "OK":
				fmt.Println("Utente eliminato con successo!")
			}
			comm.SetToken("")
			return

		case 2:
			Menu2(comm)
		case 3:
			return
		}
	}

}

func Menu2(comm *api.Communication) {
	for {
		var v int
		for {
			fmt.Println("\n1)Gestisci Parametri")
			fmt.Println("2)Gestisci Investimenti")
			fmt.Println("3)Indietro")
			fmt.Scan(&v)
			if v >= 1 || v <= 3 {
				break
			} else {
				fmt.Println("Scelta ERRATA! Ripeti la scelta..")
			}
		}

		switch v {
		case 1:
			MenuParameters(comm)
		case 2:
			MenuInvestments(comm)
		case 3:
			return
		}
	}
}

func MenuParameters(comm *api.Communication) {
	for {
		var v int
		fmt.Println("\nGestisci Parametri")
			for {
				fmt.Println("1)Crea Parametri")
				fmt.Println("2)Visualizza Parametri")
				fmt.Println("3)Elimina Parametri")
				fmt.Println("4)Indietro")
			fmt.Scan(&v)
			if v >= 1 || v <= 4 {
				break
			} else {
				fmt.Println("Scelta ERRATA! Ripeti la scelta..")
			}
		}

		switch v {
		case 1:
			fmt.Println("Inserisci il numero totale di investitori")
			var nInvest int
			fmt.Scan(&nInvest)
			fmt.Println("Inserisci il numero di investitori real-time")
			var NumbRtPlayers int
			fmt.Scan(&NumbRtPlayers)
			fmt.Println("Inserisci il prezzo della cpu per millicore")
			var PriceCpu float32
			fmt.Scan(&PriceCpu)
			fmt.Println("Inserisci la capacità massima che il Network Operator può ospitare")
			var HostingCapacity int64
			fmt.Scan(&HostingCapacity)
			fmt.Println("Inserisci la durata della CPU in anni")
			var DurationCpu int
			fmt.Scan(&DurationCpu)
			params := models.ParametersReq{InvestorsNumber: nInvest, NumbRtPlayers: NumbRtPlayers, PriceCpu: PriceCpu,
				HostingCapacity: HostingCapacity, DurationCpu: DurationCpu, UserId: comm.CurrentUser.ID}
			res := comm.CreateParameters(params)
			switch res["detail"] {
			case "Bad request: some values in parameters are not allowed":
				fmt.Println("Qualcuno dei valori inseriti è errato!")
			case "Token expired redo the login!":
				fmt.Println("Il token è scaduto: rifai il login!")
			default:
				fmt.Println("Parametri inseriti con successo!")
				utils.PrintMap(res)
			}
		case 2:
			res := comm.ReadParamsUser()
			utils.PrintListJson(res)

		case 3:
			res := comm.ReadParamsUser()
			fmt.Println("Scegli quale set di parametri eliminare...")
			utils.PrintListJson(res)
			for {
				var v int
				for {
					fmt.Println("\nInserisci il numero corrispondente ai parametri da eliminare")
					fmt.Scan(&v)
					if v >= 1 || v <=  len(res){
						break
					} else {
						fmt.Println("Scelta ERRATA! Ripeti la scelta..")
					}
				}
				paramsId:= res[v-1]["id"]
				tmp:=comm.DeleteParameters(int(paramsId.(float64)))

				switch tmp {
				case "Parameters not found!":
					fmt.Println("Username or Password non presenti!")
				case "Token expired redo the login!":
					fmt.Println("Il token è scaduto: rifai il login")
				case "OK":
					fmt.Println("Parametri eliminati con successo!")
				}
				break

				}
		case 4:
			return
			}

		}
	}

func MenuInvestments(comm *api.Communication) {
	for {
		var v int
		fmt.Println("\nGestisci Investimenti")
		for {
			fmt.Println("1)Crea Investimento")
			fmt.Println("2)Visualizza Investimenti")
			fmt.Println("3)Elimina Investimento")
			fmt.Println("4)Indietro")
			fmt.Scan(&v)
			if v >= 1 || v <= 4 {
				break
			} else {
				fmt.Println("Scelta ERRATA! Ripeti la scelta..")
			}
		}

		switch v {

		case 1:
			res := comm.ReadParamsUser()
			utils.PrintListJson(res)
			for {
				var v int
				for {
					fmt.Println("Scegli quale set di parametri usare per creare l'investimento...")
					fmt.Scan(&v)
					if v >= 1 || v <= len(res) {
						break
					} else {
						fmt.Println("Scelta ERRATA! Ripeti la scelta..")
					}
				}
				paramsId := res[v-1]["id"]
				for {
					var v int
					fmt.Println("\nTipo di divisione dei guadagni e pagamenti per l'investimento")
					for {
						fmt.Println("1)Fair")
						fmt.Println("2)Unfair")
						fmt.Scan(&v)
						if v >= 1 || v <= 2 {
							break
						} else {
							fmt.Println("Scelta ERRATA! Ripeti la scelta..")
						}
					}

					var fairness bool
					switch v {
					case 1:
						fairness = true
					case 2:
						fairness = false
					}

					invReq := models.InvestmentReq{Fairness: fairness, ParametersId: int(paramsId.(float64))}
					res := comm.CreateInvest(invReq)
					switch res["detail"] {

					case "Token expired redo the login!":
						fmt.Println("Il token è scaduto: rifai il login!")
					case "Parameters not found, parameters_id is wrong" :
						fmt.Println("Parametri non trovati, l'id potrebbe essere errato!")
					case "Not Found":
						fmt.Println("Parametri non trovati, l'id potrebbe essere errato!")
					default:
						fmt.Println("Investimento inserito con successo!")
						utils.PrintMap(res)
					}
						return
				}
			}

			case 2:
				res := comm.ReadInvestUser()
				utils.PrintListJson(res)

			case 3:
				res := comm.ReadInvestUser()
				fmt.Println("Scegli quale investimento eliminare...")
				utils.PrintListJson(res)
				for {
					var v int
					for {
						fmt.Println("\nInserisci il numero corrispondente all'investimento da eliminare")
						fmt.Scan(&v)
						if v >= 1 || v <=  len(res){
							break
						} else {
							fmt.Println("Scelta ERRATA! Ripeti la scelta..")
						}
					}
				investId:= res[v-1]["id"]
				tmp:=comm.DeleteInvest(int(investId.(float64)))

				switch tmp {
				case "Parameters not found!":
					fmt.Println("Username or Password non presenti!")
				case "Token expired redo the login!":
					fmt.Println("Il token è scaduto: rifai il login")
				case "OK":
					fmt.Println("Parametri eliminati con successo!")
				}
				break

			}

			case 4:
				return
			}
		}
	}