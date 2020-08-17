package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
	"net/http"
)

func postToZibal(path string, parameters string, apiKey string) string{
	var jsonStr = []byte(parameters)
	var url = "https://api.zibal.ir/" + path
	var bearer = "Bearer " + apiKey

	var req *http.Request
	var err error

	if path == "v1/wallet/list" {
		req, err = http.NewRequest("GET", url, nil)	
	} else {
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}


func platformResult(result string) string{
	switch result {
	case "1": 
		return "موفق"
	case "2": 
		return "API Key به درستی ارسال نشده است."
	case "3": 
		return "API Key صحیح نیست."
	case "4": 
		return "اجازه دسترسی به این سرویس صادر نشده‌است."
	case "5": 
		return "callbackUrl نامعتبر است."
	case "6": 
		return "یکی از فیلدهای اجباری ارسال نشده‌است. (در message نام فیلد مشخص می‌شود)"
	case "7": 
		return "IP ارسال‌کننده درخواست نامعتبر می‌باشد."
	case "8": 
		return "API Key غیرفعال است."
	case "9": 
		return "حداقل مبلغ باید 1000 ریال باشد."
	case "10": 
		return "کیف پول انتخاب شده وجود ندارد."
	case "11": 
		return "مبلغ درخواستی از موجودی کیف پول بیشتر است."
	case "12": 
		return "حداقل مبلغ تسویه 10000 ریال است."
	case "13": 
		return "تاخیر تسویه از حد مجاز اکانت شما کمتر است."
	case "14": 
		return "درخواست تسویه مورد نظر وجود ندارد."
	case "15": 
		return "این مقدار تاخیر تسویه برای حساب کاربری شما مجاز نمی‌باشد."
	case "16": 
		return "دسترسی این نوع درخواست تسویه برای کیف پول مورد نظر وجود ندارد."
	case "17": 
		return "امکان ثبت درخواست تسویه آنی برای مبالغ بیشتر از 50 میلیون تومان وجود ندارد."
	case "18": 
		return "مرچنت درگاه مورد نظر مورد یافت نشد و یا غیر فعال است."
	case "20": 
		return "نام وارد نشده‌است."
	case "21": 
		return "شماره شبای وارد شده نامعتبر است (شروع با IR و 26 کاراکتر)"
	case "22": 
		return "ذی‌نفع قبلا ثبت شده است."
	case "23": 
		return "ذی‌نفع نامعتبر است."
	case "24": 
		return "ذی‌نفع غیرفعال است."
	case "25": 
		return "امکان ویرایش ذی‌نفع فعال وجود ندارد."	
	case "26":
		return "امکان انجام ریفاند وجود ندارد"
			
	}	

	return "خطا در پرداخت"	
}