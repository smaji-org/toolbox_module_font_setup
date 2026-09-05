DATA?=./data

debug:
	make -C installer debug
	make -C uninstaller debug

release:
	go_cross_build cross_build.json installer build
	go_cross_build cross_build.json uninstaller build

package: release
	go_cross_exec cross_build.json . ./package.sh $(DATA)

clean:
	make -C installer clean
	make -C uninstaller clean

distclean: clean
	rm -rf build package
