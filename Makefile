this:
	mkdir -v ./static/css || echo
	sassc ./sass/fv_common.sass ./static/css/fv_common.css
	sassc ./sass/fv_dark.sass ./static/css/fv_dark.css
	sassc ./sass/fv_light.sass ./static/css/fv_light.css
	sassc ./sass/ghmd.sass ./static/css/ghmd.css
	sassc ./sass/ghmd_dark.sass ./static/css/ghmd_dark.css
	sassc ./sass/ghmd_light.sass ./static/css/ghmd_light.css
	sassc ./sass/fonts.scss ./static/css/fonts.css
	sassc ./sass/style.sass ./static/css/style.css
	./scripts/add_css.sh
	go build .

