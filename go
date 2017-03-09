#!/bin/bash


if [[ $1 = "ios" ]]; then
	~/King/Source/candycrushsaga/go.sh build_configs/AppStore-Facebook-QA-debug.json update -c
elif [[ $1 = "android" ]]; then
	~/King/Source/candycrushsaga/go.sh build_configs/GooglePlay-Facebook-QA-release.json build -c
	if [[ "$?" -eq "0" ]]; then
		adb uninstall com.king.candycrushsaga
		adb install projects/GooglePlay-Facebook-QA-release/android/bin/CandyCrushSaga-release-signed.apk
	fi
else
	~/King/Source/candycrushsaga/go.sh build_configs/OSX-GooglePlay-Facebook-QA.json update -c
fi
