// Ajax

function ajaxGet(requestUrl, callback) {
    $.ajax({
        type: `GET`,
        url: requestUrl,
        success: function(response) {
            callback(response)
        }
    })
}

function ajaxGetSync(requestUrl, callback) {
    $.ajax({
        type: `GET`,
        url: requestUrl,
        async: false,
        success: function(response) {
            callback(response)
        }
    })
}

function ajaxPostUpload(requestUrl, data, callback) {
    $.ajax({
        type: `POST`,
        url: requestUrl,
        contentType: false,
        processData: false,
        data: data,
        success: function(response) {
            callback(response)
        }
    })
}

function ajaxPostJson(requestUrl, data, callback) {
    $.ajax({
        type: `POST`,
        url: requestUrl,
        contentType: `application/json`,
        data: JSON.stringify(data),
        dataType: `json`,
        success: function(response) {
            callback(response)
        } 
    })
}

function ajaxPostJsonSync(requestUrl, data, callback) {
    $.ajax({
        type: `POST`,
        url: requestUrl,
        async: false,
        contentType: `application/json`,
        data: JSON.stringify(data),
        dataType: `json`,
        success: function(response) {
            callback(response)
        } 
    })
}

function ajaxPutJson(requestUrl, data, callback) {
    $.ajax({
        type: `PUT`,
        url: requestUrl,
        contentType: `application/json`,
        data: JSON.stringify(data),
        dataType: `json`,
        success: function(response) {
            callback(response)
        } 
    })
}

function ajaxDeleteJson(requestUrl, data, callback) {
    $.ajax({
        type: `DELETE`,
        url: requestUrl,
        contentType: `application/json`,
        data: JSON.stringify(data),
        dataType: `json`,
        success: function(response) {
            callback(response)
        } 
    })
}

// LocalStorage and SessionStorage

function setLocal(key, val) {
    window.localStorage.setItem(key, val)
}

function getLocal(key) {
    return window.localStorage.getItem(key)
}

function removeLocal(key) {
    window.localStorage.removeItem(key)
}

function setSession(key, val) {
    window.sessionStorage.setItem(key, val)
}

function getSession(key, val) {
    window.sessionStorage.getItem(key, val)
}

function removeSession(key) {
    window.sessionStorage.removeItem(key)
}

// Input limit

function inputsInt(val, defaultVal = ``, minVal = -999999999, maxVal = 999999999) {
    if (val) {
        let val1 = val + ``
        val1 = val1.replace(/[^0-9]/g, ``)
        if (val1 === ``) {
            return defaultVal
        }
        if (Number(val1) < minVal) {
            return minVal + ``
        } else if (Number(val1) > maxVal) {
            return maxVal + ``
        }
        return val1
    } else {
        return defaultVal
    }
}