var curentProtcol = window.location.protocol,
    curentHost = window.location.host,
    consoleStyle = {};
    consoleStyle['Defualt'] = "";
    consoleStyle['Debug-1'] = "color:ornage;font-weight:bold;";
    consoleStyle['Success'] = "color:green;font-weight:bold;";
    consoleStyle['Error'] = "color:red;font-weight:bold;";
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
      FAQs:"",
      forms:""
    },
    msg:"",
    bot:"https://console.dialogflow.com/api-client/demo/embedded/1f8aea9e-26c0-47d3-b699-234257524470",
    token: 'Basic Zm9ybXNjcmliZXJhcGk1MjM0NTo5ODcyMzQ4OTcydXNoZGZ1U0RGwqckwqc=',
    formTemplatesURL: "",
    formList:[],
    FAQs:{},
    howToIndex:0,
    menuIndex:0,
    menuType:"PC",
    showMenu: true,
    showMobileMenu: false,
    showFormTemplates: false,
    docView: false,
    docViererURLid:"",
    error:[]
  },
// you would use this to load data. for this you will not need to wory about it
  mounted: function () {
    var self = this, 
    getArticles = curentProtcol + "//" + curentHost + "/getArticles",
    getTeam = curentProtcol + "//" + curentHost + "/getTeam",
    getFAQs = curentProtcol + "//" + curentHost + "/getFAQs";
    
    axios.all([
      axios.get(getArticles),
      axios.get(getTeam),
      axios.get(getFAQs)
    ])
      .then(r => {
// handle success
//      console.log('%cSUCCESS', consoleStyle['Success']);
//      console.log(r[3]);
        self.articles = r[0].data.d;
        self.teamRoster = r[1].data.d;
        var tempFAQarr = r[2].data.d;
        tempFAQarr.map(o => (o.show = false));
        self.FAQs =tempFAQarr;
      })
      .catch(function (e) {
// handle error
        console.log('%cERROR: Mounting data', onsoleStyle['Error']);
        console.log(e);
        self.error = e;
      })
      .then(function () {
// always executed
      });        
    },
// computer is where you can create functions that can mutated that data as the view is refreshed
    computed:{
      filteredForms() {
        var filterKey = this.filterKey.forms && this.filterKey.forms.toLowerCase();
        var data = this.formList;
        if (filterKey) {
          data = data.filter(function (row) {
            return Object.keys(row).some(function (key) {
              return String(row[key]).toLowerCase().indexOf(filterKey) > -1;
            })
          })
        }
        return data;
      },
      filteredFAQs() {
        var filterKey = this.filterKey.FAQs && this.filterKey.FAQs.toLowerCase();
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
      capitalize (str) { return str.charAt(0).toUpperCase() + str.slice(1);}
    },
    methods: {      
      isMobile() {
        var check = false,
          self = this;
        (function(a){if(/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i.test(a)||/1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(a.substr(0,4))) {check = true; self.menuType = "mobile"; self.showMenu = false;console.log('***MOBILE');}})(navigator.userAgent||navigator.vendor||window.opera);        
        return check;
      },
      tab(id){
        this.msg = "";
        this.formTemplatesURL = "";
        this.showFormTemplates = false;
        this.formList.length = 0;
        this.menuIndex = id;
        this.showMobileMenu = false;
      },
      howTo(i){this.howToIndex = i;},
      togelFAQ(index){
        this.FAQs.forEach(function(o, i){
          if(i == index){ o.show = true;}
          else{o.show = false;}
        });
      },
      togelMenu(){
        var self = this;
        if(self.showMobileMenu){ self.showMobileMenu = false;}
        else{self.showMobileMenu = true;}
      },
      copyTextArea(fileName, url) { 
        navigator.clipboard.writeText(url);
        this.msg = '<em>"' + fileName + '"</em> has been copied."';
      },
      formURL(id){ return "https://docs.google.com/document/d/" + id;},
      toggelDocViewer(id){
        var self = this;
        if(self.docView){
          self.docView = false;
        }else{
          self.docView = true;
          self.docViererURLid = id;
        }
      },
      pullTemplateList(){
        var self = this;

        if(self.formTemplatesURL !== ""){
           var config = {
                method: 'post',
                url: curentProtcol + "//" + curentHost + '/drive',
                headers: { 
                  'Authorization': self.token
                },
                data: JSON.stringify({"driveUrl": self.formTemplatesURL})
            };
            
            axios(config)
            .then(function (r) {
//            console.log('%cSUCCESS', consoleStyle['Success']);
//            console.log(r);
              self.formList = r.data.files;
              self.showFormTemplates = true;
            })
            .catch(function (e) {              
              console.log('%cERROR: Form Template List Pull', consoleStyle['Error']);
              console.log(e);
            });//*/

  //          self.showFormTemplates = true;
        }
      }
    }
});

