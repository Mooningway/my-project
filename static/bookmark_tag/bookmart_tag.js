$(function() {
    let $topOperate = $(`#top-operate`)
    let $datas = $(`#datas`)
    let $editor = $(`#editor`)

     $(`#toadd`).on(`click`, () => {
        $topOperate.addClass(`hide`)
        $datas.addClass(`hide`)
        $editor.removeClass(`hide`)
     })

     $(`#editor-cancel`).on(`click`, () => {
        $topOperate.removeClass(`hide`)
        $datas.removeClass(`hide`)
        $editor.addClass(`hide`)
     })

     getBookmarkTags()
  
    function getBookmarkTags(element) {
        ajaxPostJson(`/api/bookmark/tag/page`, {}, response => {
            let bookmarks = data.data
            if (bookmarks) {
                let html = []
                $(element).html(html.join(``))
            }
        })
    }

    function addBookmarkTag(data) {
        ajaxPostJson(`/api/bookmark/tag/`, data, response => {

        })
    }

    function ediotrHandle(operate) {
        let $formId = $(`#form-id`)
        let $formName = $(`#form-name`)
        let $formTag = $(`#form-tag`)
        let $formSort = $(`#form-sort`)

        if (operate) {
            if (`reset` === operate) {
                $formId.val(``)
                $formName.val(``)
                $formTag.val(``)
                $formSort.val(0)
            }
        }
    }
})