module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    wiredep: {
      target: {
        exclude: ['bootstrap-sass'],
        // Point to the files that should be updated when
        // you run `grunt wiredep`
        src: [
          'src/ustackweb/views/layouts/default.tpl.html'
        ]
      }
    },
    sass: {
      dist: {
        files: [{
          expand: true,
          cwd: 'src/ustackweb/static/scss',
          src: ['*.scss'],
          dest: 'src/ustackweb/static/css',
          ext: '.css'
        }]
      }
    },
    watch: {
      wiredep: {
        files: 'bower.json',
        tasks: ['shell:bower', 'wiredep'],
        options: {
          debounceDelay: 250
        }
      },
      livereload: {
        files: 'src/ustackweb/**/*',
        options: {
          debounceDelay: 250,
          livereload: true
        }
      },
      static: {
        files: 'src/ustackweb/static/**/*',
        tasks: ['sass'],
        options: {
          debounceDelay: 250
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
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.registerTask('default', ['shell:bower']);
};
