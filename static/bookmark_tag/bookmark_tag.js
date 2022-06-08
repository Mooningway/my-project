const BookmarkTagApp = {
    data() {
        return {
            bookmarkTags: [],
            bookmarkTagsView: true,
            bookmarkTagEditView: false,
            queryForm: { keyword: ``, page: 1, pageSize: 10, lastPage: 1 },
            form: { rowid: ``, name: ``, description: ``, sort: 0 },
            formDisabled: true,
            formOperate: ``,
            formOperating: false,
            formResult: { msg: ``, success: false, active: false }
        }
    },
    created() {
        this.getBookmarkTags()
    },
    methods: {
        getBookmarkTags() {
            this.bookmarkTags = []
            ajaxPostJson(`/api/bookmark/tag/page`, this.queryForm, response => {
                this.bookmarkTags = response.data
                this.queryForm.page = Number(response.page)
                this.queryForm.lastPage = Number(response.lastPage)
            })
        },
        pageClick(page) {
            this.queryForm.page = Number(page)
            this.getBookmarkTags()
        },
        toEditor(formOperate, id = 0) {
            this.formOperate = formOperate
            if (formOperate === `Add` || formOperate === `Update`) {
                this.formDisabled = false
            } else if (formOperate === `Delete`) {
                this.formDisabled = true
            }

            if (Number(id) > 0) {
                ajaxGet(`/api/bookmark/tag/` + id, reponse => {
                    let data = reponse.data
                    this.form = { rowid: data.rowid, name: data.name, description: data.description, sort: data.sort }
                })
            }

            this.bookmarkTagsView = false
            this.bookmarkTagEditView = true
        },
        formSubmit() {
            if (this.formOperate) {
                if (this.formOperate === `Add` || this.formOperate === `Update`) {
                    this.formOperating = true
                    ajaxPostJson(`/api/bookmark/tag/`, this.form, response => {
                        this.formSubmitCallback(response)
                    })
                } else if (this.formOperate === `Delete`) {
                    this.formOperating = true
                    ajaxDeleteJson(`/api/bookmark/tag/` + this.form.rowid, {}, response => {
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
                    this.getBookmarkTags()
                }, 2000)
            } else {
                this.formResult.success = false
                this.formOperating = false
            }
        },
        formReset() {
            this.bookmarkTagsView = true
            this.bookmarkTagEditView = false
            this.form = { rowid: ``, name: ``, description: ``, sort: 0 }
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

Vue.createApp(BookmarkTagApp).mount(`#bookmark-tag-app`)