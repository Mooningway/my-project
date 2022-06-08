const BookmarkApp = {
    data() {
        return {
            bookmarks: [],
            bookmarkTags: [],
            bookmarksView: true,
            bookmarkEditView: false,
            queryForm: { keyword: ``, tag: ``, page: 1, pageSize: 10, lastPage: 1 },
            form: { rowid: ``, name: ``, tag: ``, link: ``, description: ``, sort: 0 },
            formDisabled: true,
            formOperate: ``,
            formOperating: false,
            formResult: { msg: ``, success: false, active: false }
        }
    },
    created() {
        this.getBookmarks()
        this.getBookmarkTags()
    },
    methods: {
        getBookmarks() {
            this.bookmarks = []
            ajaxPostJson(`/api/bookmark/page`, this.queryForm, response => {
                this.bookmarks = response.data
                this.queryForm.page = Number(response.page)
                this.queryForm.lastPage = Number(response.lastPage)
            })
        },
        getBookmarkTags() {
            this.bookmarkTags = []
            ajaxGet(`/api/bookmark/tag`, response => {
                this.bookmarkTags = response.data
            })
        },
        queryFormRest() {
            this.queryForm.keyword = ``
            this.queryForm.tag = ``
        },
        pageClick(page) {
            this.queryForm.page = Number(page)
            this.getBookmarks()
        },
        toEditor(formOperate, id = 0) {
            this.formOperate = formOperate
            if (formOperate === `Add` || formOperate === `Update`) {
                this.formDisabled = false
            } else if (formOperate === `Delete`) {
                this.formDisabled = true
            }

            if (Number(id) > 0) {
                ajaxGet(`/api/bookmark/` + id, reponse => {
                    let data = reponse.data
                    this.form = { rowid: data.rowid, name: data.name, tag: data.tag, link: data.link, description: data.description, sort: data.sort }
                })
            }

            this.bookmarksView = false
            this.bookmarkEditView = true
        },
        formSubmit() {
            if (this.formOperate) {
                if (this.formOperate === `Add` || this.formOperate === `Update`) {
                    this.formOperating = true
                    ajaxPostJson(`/api/bookmark`, this.form, response => {
                        this.formSubmitCallback(response)
                    })
                } else if (this.formOperate === `Delete`) {
                    this.formOperating = true
                    ajaxDeleteJson(`/api/bookmark/` + this.form.rowid, {}, response => {
                        this.formSubmitCallback(response)
                    })
                }
            }
        },
        formSubmitCallback(response) {
            this.formResult.msg = response.msg
            this.formResult.active = true
            if (response.code === `200`) {
                this.formResult.success = true
                setTimeout(() => {
                    this.formReset()
                    this.getBookmarks()
                    this.getBookmarkTags()
                }, 2000)
            } else {
                this.formResult.success = false
                this.formOperating = false
            }
        },
        formReset() {
            this.bookmarksView = true
            this.bookmarkEditView = false
            this.form = { rowid: ``, name: ``, tag: ``, link: ``, description: ``, sort: 0 }
            this.formDisabled = true
            this.formResult = { msg: ``, success: false, active: false }
            this.formOperating = false
        },
        sortChange() {
            let temp = inputsInt(this.form.sort, 0, 0, 99)
            this.form.sort = Number(temp)
        }
    }
}

Vue.createApp(BookmarkApp).mount(`#bookmark-app`)