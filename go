#!/bin/bash

GO_SCRIPT=~/King/Source/candycrushsaga/go.sh
OSX_CONFIG="build_configs/OSX-GooglePlay-Facebook-QA.json"
IOS_CONFIG="build_configs/AppStore-Facebook-QA-debug.json"
ANDROID_CONFIG="build_configs/GooglePlay-Facebook-QA-debug.json"
ANDROID_PACKAGE="projects/GooglePlay-Facebook-QA-debug/android/bin/CandyCrushSaga-debug-signed.apk"
ANDROID_BUNDLE_ID="com.king.candycrushsaga"

function update {
	CONFIG_FILE="$1"
	shift
	echo "Updating using config file: ${CONFIG_FILE}"
	${GO_SCRIPT} ${CONFIG_FILE} update -c "$@"
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
		r|resource)
		echo "Updating resources"
		${GO_SCRIPT} ${ANDROID_CONFIG} resources
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
	c|check)
	tools/scripts/build.py --policies --targets ${OSX_CONFIG}
	checkbranch
	;;
	*)
	shift
	update ${OSX_CONFIG} "$@"
	;;
esac
