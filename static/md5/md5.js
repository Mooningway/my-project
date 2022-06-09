function md5(data, callback) {
    ajaxPostJson(`/api/md5`, {source: data}, response => {
        
        callback(response.data)
    })
}

const MD5App =  {
    data() {
        return {
            source: ``,
            result: ``
        }
    },
    methods: {
        encrypteClick() {
            md5(this.source, response => {
                this.result = response
            })
        }
    }
}

Vue.createApp(MD5App).mount(`#md5-app`)