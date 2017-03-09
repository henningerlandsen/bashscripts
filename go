#!/bin/bash

GO_SCRIPT=~/King/Source/candycrushsaga/go.sh
OSX_CONFIG="build_configs/OSX-GooglePlay-Facebook-QA.json"
IOS_CONFIG="build_configs/AppStore-Facebook-QA-debug.json"
ANDROID_CONFIG="build_configs/GooglePlay-Facebook-QA-release.json"
ANDROID_PACKAGE="projects/GooglePlay-Facebook-QA-release/android/bin/CandyCrushSaga-release-signed.apk"
ANDROID_BUNDLE_ID="com.king.candycrushsaga"

function update {
	echo "Updating using config file: $1"
	${GO_SCRIPT} $1 update -c
}

function android {
	case $1 in
		i|install)
		echo "Uninstalling existing bundle..."
		adb uninstall ${ANDROID_BUNDLE_ID}
		echo "Installing..."
		adb install ${ANDROID_PACKAGE}
		;;
		u|update)
		update ${ANDROID_CONFIG}
		;;
		b|build)
		echo "Building"
		pushd projects/GooglePlay-Facebook-QA-release
		make -j 4
		popd
		;;
	esac
}


case $1 in
	i|ios)
	update ${IOS_CONFIG}
	;;
	a|android)
	android $2
	;;
	*)
	update ${OSX_CONFIG}
	;;
esac
