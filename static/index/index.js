const IndexApp = {
    data() {
        return {
            searchEngines: [],
            searchEngineIndex: 0,
            goUrl: ``,
            goKeyword: ``,
            bookmarkTags: [],
            bookmarkTagValue: ``,
            bookmarks: [],
        }
    },
    created() {
        this.getSearchEngines()
        this.getBookmarkTags()
    },
    methods: {
        getSearchEngines() {
            this.searchEngines = []
            ajaxGet(`/api/searchengine/all`, response => {
                this.searchEngines = response.data
                this.goUrl = response.data[this.searchEngineIndex].url
            })
        },
        getBookmarkTags() {
            this.bookmarkTags = []
            ajaxGet(`/api/bookmark/tag`, response => {
                this.bookmarkTags = response.data
            })

        },
        getBookmarks(tag) {
            this.bookmarks = []
            ajaxGet(`/api/bookmark/bytag/` + tag, response => {
                this.bookmarks = response.data
            })
        },
        keywordChange() {
            let se = this.searchEngines[this.searchEngineIndex]

            let goTo = []
            goTo.push(se.url)
            if (this.goKeyword && this.goKeyword.length > 0) {
                goTo.push(se.search)
                goTo.push(this.goKeyword)
            }
            this.goUrl = goTo.join(``)
        },
        tagClick(tag) {
            this.bookmarkTagValue = tag
            this.getBookmarks(tag)
        }
    }
}

Vue.createApp(IndexApp).mount(`#index-app`)