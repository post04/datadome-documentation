var a = {
    "ttst": 19.100000005215406, // time taken to execute the last async task using performance.Now()
    "ifov": false, // if an error occured while creating elements
    "tagpu": 4.800000000745058, // time taken to do their shitty webGL stuff
    "glvd": "Google Inc. (Google)", // webgl vendor
    "glrd": "ANGLE (Google, Vulkan 1.3.0 (SwiftShader Device (Subzero) (0x0000C0DE)), SwiftShader driver)", // webgl render driver
    "hc": 32, // window.navigator.hardwareConcurrency
    "br_oh": 1032, // window.navigator.outerHeight
    "br_ow": 1920, // window.navigator.outerWidth
    "ua": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 OPR/93.0.0.0", // user agent
    "wbd": false, // !!window.webdriver (always false unless you wanna get flagged lol)
    "wdif": false, // !!this._iframeRef.contentWindow.navigator.webdriver aka same shit as last one
    "wdifrm": false, // this._iframeRef.contentWindow === window || this._iframeRef.contentWindow.setTimeout === window.setTimeout
    "npmtm": false, // if all post messages sent to the _iframeRef, what it's sending is navigator.plugins and stuff like that, so static false is fine
    "br_h": 931, // Math.max(document.documentElement.clientHeight, window.innerHeight || 0) https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth/dimensions-client.png
    "br_w": 1879, // Math.max(document.documentElement.clientWidth, window.innerWidth || 0) https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth/dimensions-client.png
    // I want to make this more clear, this .match returns an array with 1 element so you're not getting .length of the cookie, you're getting .length of the array. So if you have a previous cookie, it should be 1 theoretically
    "nddc": 0, // (document.cookie.match(/datadome=/g) || []).length
    "rs_h": 1080, // window.screen.height
    "rs_w": 1920, // window.screen.width
    "rs_cd": 24, // window.screen.colordepth
    "phe": false, // !(!window.callPhantom && !window._phantom.) always false
    "nm": false, // !!window.__nightmare always false
    "jsf": false, // hard code is fine, checks a bunch of stuff (!Function.prototype.bind || Function.prototype.bind.toString().replace(/bind/g, 'Error') != Error.toString() && void 0 === window.Prototype) then sets to false 
    "lg": "en-US", // navigator.language || navigator.userLanguage || navigator.browserLanguage || navigator.systemLanguage || ''
    "pr": 1, // window.devicePixelRatio
    "ars_h": 1032, // screen.availHeight
    "ars_w": 1920, // screen.availWidth
    "tz": 300, // new Date().getTimezoneOffset()
    "str_ss": true, // !!window.sessionStorage always true
    "str_ls": true, // !!localStorage always true
    "str_idb": true, // !!indexedDB always true
    "str_odb": true, // !!openDatabase always true
    "plgod": false, // !!Object.getOwnPropertyDescriptor(navigator, "plugins")
    "plg": 5, // navigator.plugins.length
    "plgne": true, // navigator.plugins[0].name === navigator.plugins[0][0].enabledPlugin.name always true I guess lol
    "plgre": true, // reverse of previous
    "plgof": false, // navigator.plugins.item(859523698994125) === navigator.plugins[0] always false, idk what world this would be true LOL
    "plggt": false, // Object.getOwnPropertyDescriptor(navigator.__proto__, "plugins").get.toString().indexOf("return") > -1 always false I guess
    "pltod": false, // !!Object.getOwnPropertyDescriptor(navigator, "platform")
    "hcovdr": false, // checks if Object.setPrototypeOf(Object.getOwnPropertyDescriptor(navigator.__proto__, 'hardwareConcurrency').get, Object.getOwnPropertyDescriptor.(navigator.__proto__, 'hardwareConcurrency').get) results in an error, if so, return false. Aka this is always false
    "hcovdr2": false, // same for this 
    "plovdr": false, // same for this except platform instead of hardwareConcurrency
    "plovdr2": false, // same for this except platform instead of hardwareConcurrency
    "ftsovdr": false, // same except instead of Object.getOwnPropertyDescriptor it's Function.prototype.toString
    "ftsovdr2": false, // same except instead of Object.getOwnPropertyDescriptor it's Function.prototype.toString
    "lb": false, // always false, just checks for inconsistencies between the length of eval.toString().length and your user agent
    "eva": 33, // length of eval.toString().length. 33 should be chrome, opera, or other. 39 should be internet explorer or other. 37 should be safari, firefox, or other.
    "lo": false, // always false, checks for inconsistencies between you user agent, cpu, and platform. Don't feel like documenting all that it's just false
    "ts_mtp": 0, // navigator.maxTouchPoints || navigator.msMaxTouchPoints || 0
    "ts_tec": false, // always false, means you didn't error when creating an event "TouchEvent"
    "ts_tsa": false, // 'ontouchstart'in window so I guess always false
    "vnd": "Google Inc.", // window.navigator.vendor
    "bid": "NA", // window.navigator.buildID || 'NA'
    "mmt": "application/pdf,text/pdf", // all window.navifator.mimTypes type's joined by nothing so [].join()
    "plu": "PDF Viewer,Chrome PDF Viewer,Chromium PDF Viewer,Microsoft Edge PDF Viewer,WebKit built-in PDF", // same except plugins
    "hdn": false, // !!document.hidden always false
    "awe": false, // !!document.awesomium always false
    "geb": false, // !!document.gib always false
    "dat": false, // 'docAutomation'in window || 'domAutomationController'in window always false
    "med": "defined", // window.navigator.mediaDevices ? 'defined' : 'NA'
    "aco": "NA", // is oog supported by an audio element, later it will do the same check but for MediaSource || WebKidMediaSource
    "acmp": "NA", // mpeg
    "acw": "NA", // wav
    "acma": "NA", // x-m4a
    "acaa": "NA", // aac
    "ac3": "NA", // 3gpp
    "acf": "NA", // flac
    "acmp4": "NA", // mp4
    "acmp3": "NA", // mp3
    "acwm": "NA", // webm
    "acots": "NA", // oog
    "acmpts": "NA", // mpeg
    "acwts": "NA",  // wav
    "acmats": "NA", // x-m4a
    "acaats": "NA", // aac
    "ac3ts": "NA", // 3gpp
    "acfts": "NA", // flac
    "acmp4ts": "NA", // mp4
    "acmp3ts": "NA", // mp4
    "acwmts": "NA", // webm
    "vco": "probably", // ogg | same as last except video element and video formats
    "vcots": false, // ogg
    "vch": "probably", // mp4
    "vchts": true, // mp4
    "vcw": "probably", // webm
    "vcwts": true, // webm
    "vc3": "maybe", // 3gpp
    "vc3ts": false, // 3gpp
    "vcmp": "", // mpeg
    "vcmpts": false, // mpeg
    "vcq": "", // quicktime
    "vcqts": false, // quicktime
    "vc1": "probably", // mp4
    "vc1ts": true, // mp4
    "dvm": 8, // navigator.deviceMemory || 'NA'
    "sqt": false, // always false, window.external && window.external.toString && window.external.toString().indexOf('Sequentum') > -1
    "so": "landscape-primary", // window.screen.orientation.type
    "wdw": true, // checks if ur on a chromeium based browser, if you are, set to true
    "cokys": "bG9hZFRpbWVzY3Npc2VhcmNoYXBwL=", // if wdw is true, add the names on every function in window.chrome to a string then base64 encode it and add the string "L=" to the end
    "ecpc": false, // always false, checks if it's an electron proccess via !!window.process
    "lgs": true, // '' !== navigator.languages
    "lgsod": false, // !!Object.getOwnPropertyDescriptor(navigator, 'languages')
    "psn": true, // !!window.PermissionStatus && window.PermissionStatus.prototype.hasOwnProperty('name')
    "edp": true, // !!window.EyeDropper
    "addt": true, // !!window.AudioData
    "wsdc": true, // !!window.WritableStreamDefaultController
    "ccsr": true, // !!window.CSSCounterStyleRule
    "nuad": true, // !!window.NavigatorUAData
    "bcda": false, // !!window.BarcodeDetector
    "idn": true, // !(!window.Intl || !Intl.DisplayNames)
    "capi": false, // !!(window.navigator && window.navigator.contacts && window.navigator.ContactsManager)
    "svde": false, // !!window.SVGDiscardElement
    "vpbq": true, // !!(window.HTMLVideoElement && window.HTMLVideoElement.prototype && window.HTMLVideoElement.prototype.getVideoPlaybackQuality)
    "ucdv": false, // 'undefined' != typeof objectToInspect && null === objectToInspect && 'undefined != typeof result && !result
    "spwn": false, // !!window.spawn
    "emt": false, // !!window.emit
    "bfr": false, // !!window.Buffer
    "dbov": false, // !!('' + window.console.debug).match(/[\)\( ]{3}[>= ]{3}\{\n[ r]{9}etu[n r]{3}n[lu]{3}/)
    "prm": true, // permissions check, always true ig
    "tzp": "America/New_York", // Intl.DateTimeFormat().resolvedOptions().timeZone || 'NA'
    "cvs": true, // !(!canvasElement.getContext || !canvasElement.getContext('2d'))
    "usb": "defined", // window.navigator.usb ? 'defined' : 'NA'
    "jset": 1671764364, // time in seconds Math.floor(Date.now() / 1000)
    // <!> NOTE <!>
    // EVERY VALUE AFTER THIS SEEMS TO ONLY BE PRESENT IN jsType: le not ch
    // the flow is ch -> le with ch cookie
    "cfpfe": "RXJyb3I6IENhbm5vdCByZWFkIHByb3BlcnRpZXMgb2YgbnVsbCAocmVhZGluZyAndG9TdHJpbmcnKQ==", // base64 encoded error: "Error: Cannot read properties of null (reading 'toString')"
    // this error contains a line number and is 100% triggered on purpose, if I was making an actual API, I would NOT hard code this for every website, sounds like a terrible idea truly.
    // then again it's datadome, doubt they even care lol
    "stcfp": "MzUzYTZiY2Y2NDQ5MzM5ZTQ3Nzc1ZGEwMzI3YWM3Lm1pbi5qczo2OjMxMDE3KQogICAgYXQgZS5leHBvcnRzIChodHRwczovL2Fzc2V0cy5hZG9iZWR0bS5jb20vbGF1bmNoLUVOMmUzNTNhNmJjZjY0NDkzMzllNDc3NzVkYTAzMjdhYzcubWluLmpzOjY6MTg3MjEp", // base64 encoded error: '353a6bcf6449339e47775da0327ac7.min.js:6:31017)\n    at e.exports (https://assets.adobedtm.com/launch-EN2e353a6bcf6449339e47775da0327ac7.min.js:6:18721)'
    "dcok": ".footlocker.com", // window.location.hostname
    "mp_cx": 764, // all of this is this.getMousePosition, cx is mouseMovement.clientX
    "mp_cy": 268, // mouseMovement.clientY
    "mp_tr": true, // mouseMovement.isTrusted
    "mp_mx": 0, // mouseMovement.movementX (if it's moving positive it's 1, negative it's -1, neither it's 0 I guess)
    "mp_my": -1, // mouseMovement.movementY
    "mp_sx": 940, // mouseMovement.screenX
    "mp_sy": 481, // mouseMovement.screenY
    "mm_md": 106, // This one is much different, math below
    /*
        (assume that w is set to a previous mouseMovement event)
        z = {
            clientX: mouseMovement.clientX,
            clientY: mouseMovement.clientY
        }
        y = Math.Sqrt((z.clientX - w.clientX) * (z.clientX - w.clientX) + (z.clientY - w.clientY) * (z.clientY - w.clientY))
        if w && (!mm_hd || y > mm_hd) {
            mm_hd = parseInt(y) (so Math.floor I guess is acceptable here in golang if you'd rather)
            w = z
        }
        if !w {
            w = z
        }
        
    */
    "es_sigmdn": 0.0000012567612990260548, // this is a calculation between the amount of events in ../eventCounters.mousemove and here, I guess so this number and the events count can't be random. Math is below.
    /*
    x = timestamp of all events in pool combined by math.Log(pool[i].timeStamp)
    q = number of all events in pool
    y = math.Log(pool[i].timeStamp) * math.Log(pool[i].timeStamp)
    Math.sqrt((q * y - x * x) / q * (q - 1)) / 1000

    so what that does is takes all the events in the "pool" of events and does all that math with them, every time a request is sent to the server, pool is cleared, or every 500 ms 
    
    then the payload is done with this

    array = array sorted small to large
    u = (array.length - 1) * 50 / 100
    v = Math.floor(u) + 1
    if array[v] !== 0 {
        w = u - v
        return array[v] + w * (array[v + 1] - array[v])
    }else {
        return array[v]
    }
    */
    "es_mumdn": 9.084022870046235, // timestamps / events aka x / q then the same math in the second section as es_sigmdn is done to get the number
    "es_distmdn": 85.44003745317531, // difference between x and y position of mouse between events, added as `Math.sqrt(differenceX * differenceX + differenceY * differenceY)`. Same math as before applies to get the value out of the array of values
    "es_angsmdn": -2.788602265762883, // _startAngles, math below
    /*
        D = q < 4 ? q - 1 : 3
        E = pool[D]
        F = pool[pool.length - D - 1]
        function calculateAngle(m, p, q, u) {
            var v = q - m
              , w = u - p
              , x = Math['acos'](v / Math['sqrt'](v * v + (w * w)));
            return w < 0 ? -x : x;
        }
        startAngles.push(calculateAngle(pool[0][x], pool[0][y], E[x], E[y]))
        endAngles.push(calculateAngle(pool[length-1][x], pool[length-1][y], F[x], F[y]))

        then the same math as above applies to select values out of those arrays to use in the response to the server.
    */  
    "es_angemdn": 0.3617292415199321 // _endAngles, math above
}