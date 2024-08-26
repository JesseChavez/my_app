// esbuild configuration file

const watch = process.argv.includes('--watch')
const minify = process.argv.includes('--minify')

process.env.ENKI_ENV ||= 'development'

const isDevelopment = process.env.ENKI_ENV === 'development'
const isProduction = process.env.ENKI_ENV === 'production'

const path = require('path')

const sassPlugin = require('esbuild-sass-plugin').default

const config = {
  // absWorkingDir: path.join(process.cwd(), 'frontend/src'),
  entryPoints: [
    'frontend/src/anonymous.ts', 'frontend/src/application.ts',
    'frontend/src/anonymous.scss', 'frontend/src/application.scss'
  ],
  outdir: path.join(process.cwd(), 'frontend/builds'),
  bundle: true,
  plugins: [sassPlugin()],
  publicPath: '/assets',
  // Needed to ignore errors when esbuild is resolving url in sass files (external)
  external: ['*.woff2', '*.woff', '*.ttf', '*.jpg', '*.png', '*.svg'],
  define: {
    global: 'window',
    ENKI_ENV: JSON.stringify(process.env.ENKI_ENV || 'development')
  },
  sourcemap: isDevelopment,
  minify: minify || isProduction,
  logLevel: 'info'
}

console.log(config)

// === Run esbuild ==================================
require('esbuild').context(config).then(context => {
  if (watch) {
    // Enable watch mode
    context.watch()
  } else {
    // Build once and exit if not in watch mode
    context.rebuild().then(_result => {
      context.dispose()
    })
  }
}).catch(() => process.exit(1))
