
(function(Vue){
    new Vue({
        el: 'app',
        // mounted: function(){}
        created: function(){
            console.log("hello VUE 2")
        },
        data: {
            msg: "initial message"
        }


    })
    var i = vm.$isServer()
    console.log(i)
})(Vue)