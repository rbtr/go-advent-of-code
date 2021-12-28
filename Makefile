$(DIR)/main.go:
	cp skel/* $(DIR)/

$(DIR)/input:
	touch $(DIR)/input
	 
.DEFAULT:
	mkdir -p $@
	make DIR=$@ $@/main.go $@/input
