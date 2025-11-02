// esbuild configuration file

const watch = process.argv.includes('--watch')
const minify = process.argv.includes('--minify')

process.env.APP_ENV ||= 'development'

process.env.CONTEXT_PATH ||= ''

const assetContextPath = `${process.env.CONTEXT_PATH}/assets`

const isDevelopment = process.env.APP_ENV === 'development'
const isProduction = process.env.APP_ENV === 'production'

const path = require('path')

const fs = require('fs')

const sassPlugin = require('esbuild-sass-plugin').default
const manifestPlugin = require('esbuild-plugin-manifest')

const buildDir = path.join(process.cwd(), 'tmp/assets')

const packagesPath = path.resolve(process.cwd(), 'node_modules')

const sassLoadPaths = [
  packagesPath,
]

const config = {
  // absWorkingDir: path.join(process.cwd(), 'frontend/src'),
  entryPoints: [
    'frontend/src/anonymous.ts',
    'frontend/src/application.ts',
    'frontend/src/anonymous.scss',
    'frontend/src/application.scss'
  ],
  outdir: buildDir,
  format: 'esm',
  bundle: true,
  metafile: true,
  assetNames: '[name]-[hash]',
  entryNames: '[name]-[hash]',
  loader: {
    ".woff": "file",
    ".woff2": "file",
    ".png": "file",
    ".jpeg": "file",
    ".jpg": "file",
    ".svg": "file",
  },
  plugins: [
    sassPlugin({
      // cssImports: true
      loadPaths: sassLoadPaths
    }),
    manifestPlugin({
      shortNames: true
    })
  ],
  // publicPath: '/assets',
  publicPath: assetContextPath,
  // Needed to ignore errors when esbuild is resolving url in sass and imports
  // in javascript or typescript files (external)
  external: ['*.ttf', '*.bmp'],
  define: {
    global: 'window',
    APP_ENV: JSON.stringify(process.env.APP_ENV || 'development')
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
    context.rebuild().then(result => {
      fs.writeFileSync(
        path.join(buildDir, "metafile.json"),
        JSON.stringify(result.metafile)
      )
      context.dispose()
    })
  }
}).catch(() => process.exit(1))
