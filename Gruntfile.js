module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    wiredep: {
      target: {
        // Point to the files that should be updated when
        // you run `grunt wiredep`
        src: [
          'views/layouts/*'
        ]
      }
    },
    sass: {
      dist: {
        options: {
          loadPath: 'static/bower_components/bootstrap-sass/lib'
        },
        files: [{
          expand: true,
          cwd: 'static/scss',
          src: ['*.scss'],
          dest: 'static/css',
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
        files: '**/*',
        options: {
          debounceDelay: 250,
          livereload: true
        }
      },
      static: {
        files: 'static/**/*',
        tasks: ['sass'],
        options: {
          debounceDelay: 250
        }
      },
      go: {
        files: './**/*.go',
        tasks: ['shell:fmt'],
        options: {
          debounceDelay: 250
        }
      }
    },
    shell: {
        bower: {
            command: 'bower install'
        },
        fmt: {
            command: 'go fmt ./...'
        }
    }
  });

  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-wiredep');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.registerTask('default', ['shell:bower']);
};
