import bkDialog from './src/dialog'

bkDialog.install = Vue => {
    Vue.component(bkDialog.name, bkDialog)
}

export default bkDialog
