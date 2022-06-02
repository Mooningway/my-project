const ExRateApp = {
    data() {
        return {
            title: `Exchange Rate`,
            amount: 100,
            codes: [],
            rateData: [],
            fromCode: `USD`,
            toCode: `JPY`,
            dataCode: `USD`,
            convert: { msg: ``, success: false, active: false },
            update: { msg: ``, success: false, active: false }
        }
    },
    created() {
        this.getCodes()
        this.getRateData()
    },
    methods: {
        getCodes() {
            let codesJson = getLocal(`exrate-codes`)
            if (codesJson) {
                this.codes = JSON.parse(codesJson)
            } else {
                ajaxGetSync(`/api/exrate/code`, response => {
                    codesJson = response.data
                    setLocal(`exrate-codes`, codesJson)
                    this.codes = JSON.parse(codesJson)
                })
            }
        },
        getRateData() {
            ajaxGet(`/api/exrate/ratedata`, response => {
                let rateDataJson = response.data
                this.rateData = JSON.parse(rateDataJson)
            })
        },
        convertClick() {
            let data = { amount: this.amount + ``, fromCode: this.fromCode, toCode: this.toCode }
            ajaxPostJson(`/api/exrate/convert`, data, response => {
                this.convert.msg = response.msg
                if (response.code === `200`) {
                    this.convert.success = true
                } else {
                    this.convert.success = false
                }
                this.convert.active = true
            })
        },
        updateRateDataClick() {
            ajaxPutJson(`/api/exrate/rate/` + this.dataCode, {}, response => {
                let refresh = false
                this.update.msg = response.msg
                if (response.code === `200`) {
                    this.update.success = true
                    refresh = true
                } else {
                    this.update.success = false
                }
                this.update.active = true

                if (refresh === true) {
                    setTimeout(() => {
                        this.update.active = false
                        this.getRateData()
                    }, 2000)
                } else {
                    setTimeout(() => {
                        this.update.active = false
                    }, 5000)
                }
            })
        },
        deleteRateDataClick(code) {
            ajaxDeleteJson(`/api/exrate/rate/` + code, {}, response => {
                this.getRateData()
            })
        }
    }
}

Vue.createApp(ExRateApp).mount(`#exrate-app`)