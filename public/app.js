
// https://vuejs-tips.github.io/cheatsheet/
// TODO: on email form submit, validate email length
// (1-64)@(3-255).(2-3)
// if email is valid then make post request
// POST "/api/optin" body:{email}
// https://laracasts.com/discuss/channels/vue/vue-resource-post-request-no-data
// https://laracasts.com/discuss/channels/vue/sending-values-with-vue-resource?page=1
(function(Vue){
    let vm = new Vue({
        el: '#app',
        // mounted: function(){
		// console.log("vue mounted")
	// },
        created: function(){
			console.log("vue created")
			// this.$http.get('/api/optin')
        },
        data: {
            email: "",
			didSubmitEmail: false,
			invalidEmail: false
        },
		methods:{
			sendEmail: function(){
				if(!isValidemail(this.email)){
					this.invalidEmail = true
					return
				}
				this.$http.post('/api/optin',{email: this.email})
					.then((res)=>{
						console.log("email submission result",res)
						this.didSubmitEmail = true 
						// elegant info box to user

					})
			},
			closeInfobox: function(){
				this.didSubmitEmail = false 
				this.invalidEmail = false
				this.email =""
			}
		}
    })
	// console.log("vm.$isServer ?", vm.$isServer)
	let isValidemail = function(email){
		console.log("POST requested : ", email)
		// TODO if valid 
		// (1-64)@(3-255).(2-3)
		// else elegant alert box to user
		return true 
	}
})(Vue)
