package main

import "fmt"

func main() {
	var str1, str2, str3, str4 string
	fmt.Scan(&str1, &str2, &str3)
	if str1 == "vertebrado" {
		if str2 == "ave" {
			if str3 == "carnivoro" {
				str4 = "aguia"
			} else {
				str4 = "pomba"
			}
		} else {
			if str3 == "onivoro" {
				str4 = "homem"
			} else {
				str4 = "vaca"
			}
		}
	} else {
		if str2 == "inseto" {
			if str3 == "hematofago" {
				str4 = "pulga"
			} else {
				str4 = "lagarta"
			}
		} else {
			if str3 == "hematofago" {
				str4 = "sanguessuga"
			} else {
				str4 = "minhoca"
			}
		}
	}
	fmt.Println(str4)
}
