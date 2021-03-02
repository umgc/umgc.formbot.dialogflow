new Vue({
    // The following will set where in the HTML you will use VUE
    el: '#dashboard.app',
    // Data is your data for the app. it can be preloaded or empty
    data: {
        title:"FormBot Dashboard",
        articles:[],
        teamRoster:[],
        filterKey:{
            articles:"",
            faqs:"",
        },
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
        filteredArticles: function () {
//        var sortKey = this.sortKey;
          var filterKey = this.filterKey.articles && this.filterKey.articles.toLowerCase();
//        var order = this.sortOrders[sortKey] || 1;
          var data = this.articles;
          if (filterKey) {
              data = data.filter(function (row) {
                return Object.keys(row).some(function (key) {
                  return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
                })
              })
          }
/*        if (sortKey) {
            data = data.slice().sort(function(a, b) {
              a = a[sortKey];
              b = b[sortKey];
              return (a === b ? 0 : a > b ? 1 : -1) * order;
            });
          }//*/
          return data;
        }
    },    
    filters: {
        capitalize: function (str) {
        return str.charAt(0).toUpperCase() + str.slice(1)
        }
    },
    // methods are functions that react to user action
    methods: {
        tab(id){
            this.menuIndex = id;
        }
    }
});

