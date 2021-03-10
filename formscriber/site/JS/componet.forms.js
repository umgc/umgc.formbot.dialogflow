Vue.component('forms', {
  props: {
    list: Array,
    filterKey: String
  },
  computed: { 
    filteredForms: function () {
      var filterKey = this.filterKey && this.filterKey.toLowerCase();
      var data = this.list
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
  methods:{
    readForm(form){
      this.$emit('openformreader',form.id);
    },
    editForm(form){
      this.$emit('openformeditor',form.id);
    },
    copyTextArea(txt) {
        navigator.clipboard.writeText(txt);
    },
    test(){
/*      axios.post('/user', {
        "driveUrl": "https://www.googleapis.com/drive/v3/files?docid=1LJtz6CK_TA696h7j1oJwLyklpGDG9dTt"
       })
      .then(function (response) {
        console.log('**YAY**');
        console.log(response);
      })
      .catch(function (error) {
        console.log('**OPPS**');
        console.log(error);
      });//*/
      
      var data = JSON.stringify({"driveUrl":"https://drive.google.com/drive/u/0/folders/1LJtz6CK_TA696h7j1oJwLyklpGDG9dTt?ths=true"});
       
    var config = {
        method: 'post',
        url: 'https://www.formscriber.com/drive',
        headers: { 
   //       'Authorization': 'Basic Zm9ybXNjcmliZXJhcGk1MjM0NTo5ODcyMzQ4OTcydXNoZGZ1U0RGwqckwqc=', 
          'Content-Type': 'application/json'
        },
        data : data
      };
       
      axios(config)
      .then(function (response) {
        console.log(JSON.stringify(response.data));
      })
      .catch(function (error) {
        console.log(error);
      });//*/
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  /* TODO
   * - Clean up layout
   * - Add toggel to show and hide the iframe
   */
  template: `  
  <article id="docs" class="list" style="float: left; width: 300px; height: 400px;">
    <input placeholder="Enter URL to Template Directory in Google Drive"></input>
    <ul>
      <li v-for="form in filteredForms">{{form.formName}}
      <div class="btn" @click="test()">Test</div>  
      <div class="btn" @click="copyTextArea(form.URLl)">URL</div>
        <iframe src="https://docs.google.com/document/d/e/2PACX-1vSfv0ChJElCQbG0asDohdzZ90KetfqRf6jv7D3Vd8VHn3R5o5dHBgxqgkesGtQ3fnHvIdqrl8V-GcrJ/pub?embedded=true"></iframe>
      </li>
    </ul>
  </article>
      `
    })