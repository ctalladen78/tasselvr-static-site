
(function(Vue){
            console.log("hello VUE 1 from main app.js")
    new Vue({
        el: 'body',
        created: function(){
            console.log("hello VUE 2")
        }

    })
})(Vue)