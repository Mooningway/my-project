const Base64App = {
    data() {
        return {
            source: ``,
            result: ``
        }
    },
    methods: {
        encode() {
            ajaxPostJson(`/api/base64/encode`, {source: this.source}, response => {
                this.result = response.data
            })
        },
        decode() {
            ajaxPostJson(`/api/base64/decode`, {source: this.result}, response => {
                this.source = response.data
            })
        },
        fileUpload() {
            let file = this.$refs.fileRef.files[0]
            let data = new FormData()
            data.append("file", file)
            ajaxPostUpload(`/api/base64/image`, data, response => {
                this.result = response.data
            })
        },
        resetClick() {
            this.source = ``
            this.result = ``
            this.$refs.fileRef.value = ``
        }
    }
}

Vue.createApp(Base64App).mount(`#base64-app`)