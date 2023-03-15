package com.CCPayment.constant;

// dev http://74ab25e1merchant.cwallet.com/ccpayment/v1
// produce https://admin.ccpayment.com/ccpayment/v1
public class Config {
    public static final String ApiUrl = "https://admin.ccpayment.com/ccpayment/v1";
    public static final String HeaderAppId = "Appid";
    public static final String HeaderTimestamp = "Timestamp";
    public static final String HeaderSign = "Sign";
    public static final long expireDuration = 300L;

    //modify
    public static final String AppId = "202302010636261620672405236006912";
    //modify
    public static final String AppSecrete = "62fbff1f796c42c50bb44d4d3d065390";
}
