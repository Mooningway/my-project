$(function() {
    // Some elements
    let $amount = $(`#amount`)
    let $codes = $(`#codes`)
    let $fromCode = $(`#fromCode`)
    let $toCodes = $(`#toCodes`)

    let $convertResult = $(`#exrate-convert-result`);

    $amount.val(100)

    // Get codes and initialize
    let codeArray = []
    let codesJson = getLocal(`exrate-codes`)
    if (codesJson) {
        codeArray = JSON.parse(codesJson)
    } else {
        ajaxGetSync(`/exrate/code`, response => {
            codesJson = response.data
            setLocal(`exrate-codes`, codesJson)
            codeArray = JSON.parse(codesJson)
        })
    }

    let codeHtml = []
    codeArray.forEach(function(c, ci) {
        codeHtml.push(`<option value="` + c.code + `">` + c.code + ` - ` + c.name + `</option>`)
    })

    $codes.html(codeHtml.join(``))
    $fromCode.html(codeHtml.join(``))
    $toCodes.html(codeHtml.join(``))

    $codes.val(`USD`)
    $fromCode.val(`USD`)
    $toCodes.val(`JPY`)

    $amount.change(() => {
        let val = $amount.val()
        val = inputsInt(val)
        $amount.val(val)
    })

    // Convert
    $(`#convert`).click(() => {
        ajaxPostJson(`/exrate/convert`, {
            amount: $amount.val(), fromCode: $fromCode.val(), toCode: $toCodes.val()
        }, response => {
            if (response.code === `200`) {
                $convertResult.html(`<div class="exrate-convert-result msg msg-green">` + response.msg + `</div>`)
            } else {
                $convertResult.html(`<div class="exrate-convert-result msg msg-red">` + response.msg + `</div>`)
            }
        })
    })

    // Rate data manager
    let $exratePullResult = $(`#exrate-pull-result`)
    $(`#pullRate`).click(() => {
        let code = $codes.val()
        ajaxPutJson(`/exrate/rate/` + code, {}, response => {
            if (response.code === `200`) {
                $exratePullResult.html(`<span class="msg msg-green">` + response.msg + `</span>`)

                setTimeout(() => {
                    loadRatesData()
                    $exratePullResult.html(``)
                }, 2000)
            } else {
                $exratePullResult.html(`<span class="msg msg-red">` + response.msg + `</span>`)

                setTimeout(() => {
                    $exratePullResult.html(``)
                }, 5000)
            }
        })
    })

    loadRatesData()

    function loadRatesData() {
        ajaxGet(`/exrate/ratedata`, response => {
            let html = []
            let rateJson = response.data
            let rateArray = JSON.parse(rateJson)

            rateArray.forEach(function(r, ri) {
                if (r.code === `USD`) {
                    html.push(`<span class="msg-blue exrate-rate-data-item">` + r.code + `_` + r.data_string + `</span>`)
                } else {
                    html.push(`<span class="msg-blue exrate-rate-data-item">` + r.code + `_` + r.data_string + `<span class="exrate-rate-data-delete" data-code="` + r.code + `">×</span></span>`)
                }
            })

            let $exrateRateData = $(`#exrate-rate-data`)
            $exrateRateData.html(html.join(``))

            let deletes = $exrateRateData.find(`.exrate-rate-data-delete`)
            $(deletes).on(`click`, function() {
                let code = $(this).data(`code`)
                ajaxDeleteJson(`exrate/rate/` + code, {}, response => {
                    loadRatesData()
                })
                
            })
        })
    }
})