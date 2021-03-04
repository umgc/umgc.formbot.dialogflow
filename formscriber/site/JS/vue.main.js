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
        FAQs:{},
        menuIndex:0,
		error:[]
    },
    // you would use this to load data. for this you will not need to wory about it
    mounted: function () {
        var self = this, 
            test = "http://localhost:8080/VehiclesOfInterestWebServices/webresources/model.reasonforinterest",
            getArticles = "http://localhost:8081/getArticles",
            getTeam = "http://localhost:8081/getTeam",
            getFAQs = "http://localhost:8081/getFAQs";
        
            axios.all([
                axios.get(getArticles),
                axios.get(getTeam),
                axios.get(getFAQs),
                axios.get(test)
              ])
              .then(r => {
                // handle success
                console.log(r[3]);
                self.articles = r[0].data.d;
                self.teamRoster = r[1].data.d;
                var tempFAQarr = r[2].data.d;
                tempFAQarr.map(o => (o.show = false));
                self.FAQs =tempFAQarr;
            })
            .catch(function (e) {
                // handle error
                console.log(e);
                self.error = e;
            })
            .then(function () {
                // always executed
            });
        
    },
    // computer is where you can create functions that can mutated that data as the view is refreshed
    computed:{
        filteredArticles() {
//        The following is a filter feature. it takes the filter key inputed into a serach box and filters ALL values in each key of a object
          var filterKey = this.filterKey.articles && this.filterKey.articles.toLowerCase();
          var data = this.articles;
          if (filterKey) {
              data = data.filter(function (row) {
                return Object.keys(row).some(function (key) {
                  return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
                })
              })
          }
//      The following is a sort feature that will sort a coulmn within a row
//        var order = this.sortOrders[sortKey] || 1;
//        var sortKey = this.sortKey;
/*        if (sortKey) {
            data = data.slice().sort(function(a, b) {
              a = a[sortKey];
              b = b[sortKey];
              return (a === b ? 0 : a > b ? 1 : -1) * order;
            });
          }//*/
          return data;
        },
        filteredFAQs(){
//      TODO: create a filter of FAQs. Might be best to just copy filteredArticles
          var filterKey = this.filterKey.articles && this.filterKey.articles.toLowerCase();
          var data = this.FAQs;
          if (filterKey) {
              data = data.filter(function (row) {
                return Object.keys(row).some(function (key) {
                  return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
                })
              })
          }
        
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

