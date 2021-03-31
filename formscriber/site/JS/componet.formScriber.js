Vue.component('form-scriber', {
  props: {
    link: String
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
    }
  },
  filters: {
    capitalize: function (str) {
      return str.charAt(0).toUpperCase() + str.slice(1)
    }
  },
  template: `  
  <article id="bot" style="float: left;">
    <iframe
      allow="microphone;"
      :src="link">
    </iframe>
    
  </article>
      `
    })