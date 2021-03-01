new Vue({
    // The following will set where in the HTML you will use VUE
    el: '#dashboard.app',
    // Data is your data for the app. it can be preloaded or empty
    data: {
        title:"FormBot Dashboard",
        articles:[],
        faqs:{},
		error:[]
    },
    // you would use this to load data. for this you will not need to wory about it
    mounted: function () {
        var self = this, 
            url = "http://localhost:8080/getArticles";
        axios.get(url)
            .then(function (r) {
                // handle success
                console.log(r.data);
                self.articles = r.data.d;
            })
            .catch(function (e) {
                // handle error
                console.log(e);
            })
            .then(function () {
                // always executed
            });
        
    },
    // computer is where you can create functions that can mutated that data as the view is refreshed
    computed:{
    },
    // methods are functions that react to user action
    methods: {
    }
});

