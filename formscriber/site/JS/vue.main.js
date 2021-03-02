new Vue({
    // The following will set where in the HTML you will use VUE
    el: '#dashboard.app',
    // Data is your data for the app. it can be preloaded or empty
    data: {
        title:"FormBot Dashboard",
        articles:[],
        teamRoster:[],
        faqs:{},
        menuIndex:0,
		error:[]
    },
    // you would use this to load data. for this you will not need to wory about it
    mounted: function () {
        var self = this, 
            getArticles = "http://localhost:8081/getArticles",
            getTeam = "http://localhost:8081/getTeam";
        
            axios.all([
                axios.get(getArticles),
                axios.get(getTeam)
              ])
              .then(r => {
                // handle success
                console.log(r[1].data.d);
                self.articles = r[0].data.d;
                self.teamRoster = r[1].data.d;
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
        tab(id){
            this.menuIndex = id;
        }
    }
});

