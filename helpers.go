package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// GetIndexOf will get the index of a string in an array
// this is specific ONLY to the covid tracker and needs to be reworked
// but its 2AM so ill do it later. TODO
func GetIndexOf(a []summary, x string) int {
	if len(x) == 2 {
		for i, n := range a {
			if x == strings.ToLower(n.Code) {
				return i
			}
		}
	}

	for i, n := range a {
		if x == strings.ToLower(n.Name) {
			return i
		}
	}
	return 0 // TODO: make this something better
}

// IsInArray returns true if item is inside the array being checked
func IsInArray(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	// i feel like this should probably be reflect.Array instead of Slice but
	// i'm not sure it will make a big difference. Slice just supresses this error
	// so maybe this check isn't even needed? idk
	if arr.Kind() != reflect.Slice {
		fmt.Println("invalid data type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

//UserDetails lets us quickly access User info
func UserDetails(s *discordgo.Session, userID string) (user *discordgo.User, err error) {
	user, err = s.User(userID)
	if err != nil {
		fmt.Println("error:", err)
	}
	return
}

//GetAvatarURL returns the URL to the user's avatar
func GetAvatarURL() {
	return
}

//GetAvatarImage returns a type image.Image of the user's avatar
func GetAvatarImage(s *discordgo.Session, userID string) (img image.Image, err error) {
	user, err := s.User(userID)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	return s.UserAvatarDecode(user)
}

//SaveImage ...
func SaveImage(img image.Image, pname, fname string) (err error) {
	fpath := path.Join(pname, fname)

	f, err := os.Create(fpath)
	if err != nil {
		fmt.Println("failed to create path:", err)
		return
	}

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
	if err != nil {
		fmt.Println("failed to encode as jpeg:", err)
		return
	}

	f.Close()
	return nil
}

// StrArrayToInt will iterate over an array of type string, and convert it to type int
// this could probably be replaced with smarter code wherever im using this
func StrArrayToInt(a []string) (ia []int, err error) {
	for _, v := range a {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("error converting string array to int:", err)
		}
		ia = append(ia, i)
	}
	return ia, nil
}
