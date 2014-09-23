module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    wiredep: {
      target: {

        // Point to the files that should be updated when
        // you run `grunt wiredep`
        src: [
          'src/ustackweb/views/layouts/default.tpl.html'
        ]
      }
    },
    watch: {
      wiredep: {
        files: 'bower.json',
        tasks: ['shell:bower', 'wiredep'],
        options: {
          debounceDelay: 250,
        }
      }
    },
    shell: {
        bower: {
            command: 'bower install'
        }
    }
  });

  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-wiredep');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.registerTask('default', ['shell:bower']);
};
