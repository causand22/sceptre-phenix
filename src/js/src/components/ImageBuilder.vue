
<template>
    <div class="content">
        <!-- CREATE \ UPDATE MODAL -->
        <b-modal :active.sync="createModal.active" :on-cancel="resetCreateModal" has-modal-card>
            <div class="modal-card" style="width:40em">
                <header class="modal-card-head">
                    <p class="modal-card-title">Create a New Image Configuration</p>
                </header>
                <section class="modal-card-body">

                    <b-field label="Image Name" 
                        :type="createModal.nameErrType" 
                        :message="createModal.nameErrMsg">
                        <b-input type="text" v-model="newImageForm.name"></b-input>
                    </b-field>


                    <b-field label="OS"
                        :type="createModal.osErrType"
                        :message="createModal.osErrMsg" grouped>

                        <b-select placeholder="Select an OS" v-model="newImageForm.os">
                            <option value="linux">Linux</option>
                            <option value="windows">Windows</option>
                        </b-select>
                        <p class="control">
                            <button class="button is-light" @click="populate_os_defaults()">Populate OS Defaults</button>
                        </p>
                    </b-field>
                    <!-- Linux Options modal-->
                    <b-collapse class="card" animation="slide" :open="false" v-if="newImageForm.os == 'linux'">
                        <template #trigger="props">
                            <div class="card-header" role="button">
                                <p class="card-header-title">Linux Options</p>
                                <a class="card-header-icon">
                                <b-icon type="is-dark" size="is-small" :icon="props.open ? 'chevron-down' : 'chevron-up'">
                                </b-icon>
                                </a>
                            </div>
                        </template>
                        <div class="card-content">
                            <div class="content">
                                <b-field label="Variant">
                                    <b-input type="text" v-model="newImageForm.variant"></b-input>
                                </b-field>
                                <b-field label="Release">
                                    <b-input type="text" v-model="newImageForm.release"></b-input>
                                </b-field>
                                <b-field label="Mirror">
                                    <b-input type="text" v-model="newImageForm.mirror"></b-input>
                                </b-field>
                                <b-field label="Install Media">
                                    <b-input type="text" v-model="newImageForm.install_media"></b-input>
                                </b-field>
                                <b-field label="Packages">
                                    <b-taginput
                                        v-model="newImageForm.packages"
                                        placeholder="Add a package"
                                        type="is-light">
                                    </b-taginput>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>
                    <!-- Windows modal options -->
                    <b-collapse class="card" animation="slide" :open="false" v-if="newImageForm.os == 'windows'">
                        <template #trigger="props">
                            <div class="card-header" role="button">
                                <p class="card-header-title">Windows Options</p>
                                <a class="card-header-icon">
                                <b-icon type="is-dark" size="is-small" :icon="props.open ? 'chevron-down' : 'chevron-up'">
                                </b-icon>
                                </a>
                            </div>
                        </template>
                        <div class="card-content">
                            <div class="content">


                                <b-field label="Install Media ISO" 
                                    :type="createModal.winISOErrType"
                                    :message="createModal.winISOErrError"
                                    grouped>
                                    <b-input type="text" v-model="newImageForm.install_media"></b-input>
                                    <p class="control">
                                        <b-button class="button is-light" 
                                            @click="populate_windows_edition">Update Edition from ISO</b-button>
                                    </p>
                                </b-field>
                                <b-field label="Edition"
                                    :type="createModal.winEditionErrType" 
                                    :message="createModal.winEditionErrError">
                                    <b-select placeholder="select an install media..."
                                        :disabled="!validSelectWindowsEdition()"
                                        v-model="newImageForm.edition">
                                        <option v-for="( name, index ) in createModal.winEditionOptions" 
                                        :key="index" 
                                        :value="name"
                                        >
                                            {{ name }}
                                        </option>                                    
                                    </b-select>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>


                    <b-field label="Format">
                        <!-- <b-input type="text" v-model="newImageForm.format"></b-input> -->
                        <b-select placeholder="Select an image format type" v-model="newImageForm.format">
                            <option value="qcow2">.qcow2</option>
                            <option value="raw">.raw</option>
                        </b-select>
                    </b-field>

                    <b-field label="Size">
                        <b-input type="text" v-model="newImageForm.size"></b-input>
                    </b-field>


                    <b-field>
                        <b-checkbox v-model="newImageForm.include_protonuke" style="color:black">
                            Include Protonuke
                        </b-checkbox>
                    </b-field>
                    <b-field>
                        <b-checkbox v-model="newImageForm.include_miniccc"
                        style="color:black">
                            Include Miniccc
                        </b-checkbox>
                    </b-field>


                    <!-- <b-field label="Scripts">
                        <b-taginput
                            v-model="newImageForm.scripts"
                            placeholder="Add a package"
                            type="is-info">
                        </b-taginput>
                    </b-field> -->
                    <b-collapse class="card" animation="slide" :open="false">
                        <template #trigger="props">
                            <div class="card-header" role="button">
                                <p class="card-header-title">Advanced Options</p>
                                <a class="card-header-icon">
                                <b-icon type="is-dark" size="is-small" :icon="props.open ? 'chevron-down' : 'chevron-up'">
                                </b-icon>
                                </a>
                            </div>
                        </template>
                        <div class="card-content">
                            <div class="content">
                                <b-field label="Debootstrap Append">
                                    <b-input type="textarea" v-model="newImageForm.deb_append"></b-input>
                                    <!-- <b-tooltip label="Additional arguments to debootstrap"
                                        type="is-light"
                                        multilined>
                                        <b-icon icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip> -->
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.verbose_logs" style="color:black">Verbose Logs</b-checkbox>
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.compress" style="color:black">Compress</b-checkbox>
                                    <b-tooltip label="Compress image after creation (does not apply to raw image)" 
                                                    type="is-light" 
                                                    multilined>
                                        <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip>
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.ramdisk" style="color:black">Ram / Disk</b-checkbox>
                                    <b-tooltip label="Create a kernel/initrd pair in addition to a disk image" 
                                                    type="is-light" 
                                                    multilined>
                                        <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip>                                
                                </b-field>
                                <b-field label="Scripts">
                                    <b-taginput
                                        v-model="newImageForm.script_paths"
                                        placeholder="Add a script (include full path)">
                                    </b-taginput>
                                </b-field>
                                <b-field label="Overlays">
                                    <b-taginput
                                        v-model="newImageForm.overlays"
                                        placeholder="Add an overlay (include full path)">
                                    </b-taginput>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>
                </section>

                <footer class="modal-card-foot">
                    <button class="button is-danger mr-auto" @click="clear_create_form">Reset Form</button>
                    <button class="button is-light" 
                        :disabled="!validateCreateImage()" 
                        @click="create">Create Image Config</button>
                </footer>
                <!-- <footer v-else class="modal-card-foot buttons is-right">
                    <button class="button is-light" 
                        @click="edit">Edit Image Config</button>
                </footer> -->
            </div>
        </b-modal>

        <!-- BUILD MODAL -->
        <b-modal :active.sync="buildModal.active" :on-cancel="resetBuildModal" has-modal-card>
            <div class="modal-card" style="width:40em">
                <header class="modal-card-head">
                    <p class="modal-card-title">Build an Image</p>
                </header>
                <section class="modal-card-body">

                    <b-field label="Config Name">
                        <b-select v-model="buildModal.name">
                            <option v-for="( img, index ) in images" :key="index" :value="img.name">
                                {{ img.name }}
                            </option>
                        </b-select>
                    </b-field>

                    <b-field label="Output Directory">
                        <b-input type="text" v-model="buildModal.output"></b-input>
                    </b-field>

                    <b-field>
                        <b-checkbox v-model="buildModal.cache" style="color:black">Cache</b-checkbox>
                        <b-tooltip label="Cache rootfs as tar image" 
                                        type="is-light" 
                                        multilined>
                            <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                        </b-tooltip>
                    </b-field>
                    <b-field>
                        <b-checkbox v-model="buildModal.dryrun" style="color:black">Dry Run</b-checkbox>
                        <b-tooltip label="Do everything but call out to vmdb2" 
                                        type="is-light" 
                                        multilined>
                            <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                        </b-tooltip>
                    </b-field>

                    <b-field label="Verbosity">
                        <b-select v-model="buildModal.verbosity">
                            <option value="0">No verbose logs</option>
                            <option value="1">Enable verbose logs</option>
                            <option value="2">Enable very verbose logs</option>
                            <option value="4">Enable extra verbose logs</option>

                        </b-select>
                    </b-field>
                </section>

                <footer class="modal-card-foot buttons is-right">
                    <button class="button is-light" @click="build">Build Image</button>
                </footer>
            </div>
        </b-modal>

        <template v-if="images.length == 0">
        <section class="hero is-light is-bold is-large">
            <div class="hero-body">
            <div class="container" style="text-align: center">
                <h1 class="title">
                There are no image configurations!
                </h1>
                <!-- #TODO: add role allowed-->
                <b-button type="is-success" outlined @click="show_create_modal()">
                    Create An Image Configuration
                </b-button>
            </div>
            </div>
        </section>
        </template>
        <!-- configuration table -->
        <template v-else>
            <hr>
            <!-- Top search bar / + button -->
            <b-field position="is-right">
                <b-autocomplete v-model="searchName"
                                placeholder="Find an image config"
                                icon="search"
                                :data="filteredData"
                                @select="option => filtered = option">
                    <template slot="empty">
                        No results found
                    </template>
                </b-autocomplete>
                <p class='control'>
                    <button class='button' style="color:#686868" @click="searchName = ''">
                        <b-icon icon="window-close"></b-icon>
                    </button>
                </p>
                &nbsp; &nbsp;
                <b-tooltip 
                    v-if="roleAllowed('image', 'create')"
                    label="create a new image config" 
                    type="is-light" 
                    multilined>
                    <button class="button is-light" @click="updateImages(); show_create_modal()">
                    <b-icon icon="plus"></b-icon>
                    </button>
                </b-tooltip>
            </b-field>
            <!-- table -->
            <div style="margin-top: -1em;">
                <b-table
                    :data="filteredImages">
                    <template slot="empty">
                        <section class="section">
                        <div class="content has-text-white has-text-centered">
                            Your search turned up empty!
                        </div>
                        </section>
                    </template>
                    <b-table-column field="name" label="Name" width="200" sortable v-slot="props">
                        {{ props.row.name }}
                    </b-table-column>
                    <b-table-column field="size" label="Size" width="200" sortable v-slot="props">
                        {{ props.row.size }}
                    </b-table-column>
                    <b-table-column field="os" label="Os" width="200" sortable v-slot="props">
                        <!-- {{ props.row.os === 'linux' ? 
                            props.row.os + ' - ' + props.row.release : 
                            props.row.os }} -->
                        {{ props.row.os }}
                    </b-table-column>
                    <b-table-column field="release" label="Release" width="200" sortable v-slot="props">
                        {{ props.row.os === 'linux' ?
                            props.row.release :
                            props.row.edition }}
                    </b-table-column>
                    <b-table-column field="format" label="Format" width="200" sortable v-slot="props">
                        {{ props.row.format }}
                    </b-table-column>
                    <b-table-column label="Actions" width="150" centered v-slot="props">
                        <b-tooltip 
                            v-if="roleAllowed('image', 'build')"
                            label="Build" 
                            type="is-light">
                            <button
                                class="button is-primary is-small action"
                                :disabled="updating( props.row.status )"
                                @click="activateBuildModal(props.row.name)">
                                BUILD
                            </button>
                        </b-tooltip>
                        <!-- <b-tooltip label="Edit" type="is-light">
                            <button
                                class="button is-light is-small action"
                                :disabled="updating( props.row.status )"
                                @click="edit(props.row.name)">
                            <b-icon icon="edit"></b-icon>
                            </button>
                        </b-tooltip> -->
                        <b-tooltip label="Delete" type="is-light">
                            <button 
                                v-if="roleAllowed('images', 'delete', props.row.name)"
                                class="button is-light is-small action" 
                                :disabled="updating( props.row.status )" 
                                @click="delete_config(props.row.name)">
                            <b-icon icon="trash"></b-icon>
                            </button>
                        </b-tooltip>

                    </b-table-column>
                </b-table>
            </div>
        </template>
        <button @click="testMethod">Test</button>
    </div>
</template>

<script>
//   import EventBus from '@/event-bus'

export default {
    mounted () {
        // EventBus.$on( 'page-reload', ( route ) => {
        //     if (route.name == 'imagebuilder'){
        //         this.updateImages();
        //     }
        // })
    },
    async beforeDestroy() {
        clearInterval(this.update);
    },
    async created() {
        // this.isWaiting = false;
        // this.images = [];
        this.clear_create_form();
        this.updateImages();
    },
    computed: {
        filteredImages: function() {
            let images = this.images;

            var name_regx = new RegExp( this.searchName, 'i');
            var data = [];

            for (let i in images) {
                let img = images[i];
                if (img.name.match( name_regx )) {
                    data.push(img);
                }
            }
            return data;
        },
        filteredData() {
            let names = this.images.map (img => {return img.name;});
            return names.filter(
                option => {
                    return option
                        .toString()
                        .toLowerCase()
                        .indexOf( this.searchName.toLowerCase() ) >= 0
                }
            )
        },

    },

    methods: {
        testMethod(){
            console.log("Sending test")
            this.$http.get("image/test").then(
                response => {
                    response.json().then( state=>{
                        console.log(state)
                    })
                }
            )
        },
        // General page functions
        updateImages(){
            this.isWaiting = true
            this.$http.get('images').then(
                response => {
                    response.json().then( state => {
                        console.log(state.images)
                        this.images = state.images
                        for (let i = 0; i < this.images.length; i ++){
                            if (this.images[i].os == '') {
                                this.images[i].os = 'linux'
                            }
                        }
                    })
                }
            )
            this.isWaiting = false
        },

        show_create_modal(){
        this.createModal.type = "create"
        this.createModal.active = true 
        },

        //Create / Edit modal functionality 
        clear_create_form(){
            this.newImageForm = {
                name: null,
                cache: false,
                compress: false,
                deb_append: null,
                edition: null,
                format: null,
                include_miniccc: false,
                include_protonuke: false,
                install_media: null,
                mirror: null,
                os: null,
                overlays: [],
                packages: [],
                ramdisk: false,
                release: null,
                script_order: [],
                script_paths: [],
                scripts: {},
                size: null,
                variant: null,
                verbose_logs: false
                
            }
        },
        populate_os_defaults() {
            console.log("populating default values")
            this.$http.get(`image/create?os=${this.newImageForm.os}`).then(response => {
                response.json().then(state => {
                    this.update_create_form(state);
                    this.newImageForm.format = "qcow2"

                });

            }, err => {
                console.error(err);
            });

        },
        update_create_form(updateObject){
            for (const [key, value] of Object.entries(updateObject)) {
                if (!this.isEmpty(value)) {
                    this.newImageForm[key] = value
                }
            }
        },
        validateCreateImage() {
            if (!this.newImageForm.name) {
                return false;
            }
            if (/\s/.test(this.newImageForm.name)) {
                this.createModal.nameErrType = 'is-danger';
                this.createModal.nameErrMsg = 'experiment names cannot have a space';
                return false;
            }
            else if (/\./.test(this.newImageForm.name)) {
                this.createModal.nameErrType = 'is-danger';
                this.createModal.nameErrMsg = 'experiment names cannot have a period';
                return false;
            }
            else if (this.newImageForm.name == "create") {
                this.createModal.nameErrType = 'is-danger';
                this.createModal.nameErrMsg = 'experiment names cannot be create!';
                return false; 
            }
            else {
                this.createModal.nameErrType = null;
                this.createModal.nameErrMsg = null;
            }

            if (this.newImageForm.os == null) {
                return false
            } 

            if (this.validSelectWindowsEdition == false) {
                return false 
            }

            const sizeRegex = new RegExp('[0-9]+[KMG]')
            if (! sizeRegex.test(this.newImageForm.size)){
                return false 
            }
            return true;
        },

        validSelectWindowsEdition(){
            //skip if not using windows
            if (this.isEmpty(this.newImageForm.os) || this.newImageForm.os === 'linux') {
                console.log("False")
                return true;
            }
            if (this.isEmpty(this.newImageForm.install_media)){
                this.createModal.winISOErrType='is-danger'
                this.createModal.winISOErrError='Choose an installation file'
                return false;
            }

            if (this.createModal.enteredWinISO != this.newImageForm.install_media) {
                this.createModal.winEditionErrType = 'is-danger'
                this.createModal.winEditionErrError = 'ISO has changed since selecting an edition!'
                return false;
            }
            if (this.createModal.winEditionOptions.length == 0) {
                return false;
            }
            this.createModal.winEditionErrError = null;
            this.createModal.winEditionErrType = null;
            return true;
        },

        resetCreateModal(){
            this.clear_create_form();
        },

        create(){ 
            this.isWaiting = true;

            let form = this.newImageForm

            this.$http.post('image/create', form, { timeout: 0 } 
                ).then(
                _ => {          
                    console.log("Create Image Request sumbitted correctly")  
                    this.isWaiting = false;
                }, err => {
                    this.errorNotification(err);
                    this.isWaiting = false;
                }
                );
            this.createModal.active = false;
            this.resetCreateModal()
            this.updateImages()
        },

        populate_windows_edition(){
            console.log(this.newImageForm.install_media)
            this.$http.get(`image/edition?media=${this.newImageForm.install_media}`).then(response => {
                // console.log(response)
                response.json().then(state => {
                    // console.log(state);
                    this.createModal.winISOErrType = null,
                    this.createModal.winISOErrError = null,
                    this.createModal.winEditionOptions = state.editions

                    this.createModal.enteredWinISO = this.newImageForm.install_media

                });
            }, err => {
                this.createModal.winISOErrType = 'is-danger',
                this.createModal.winISOErrError = 'invalid ISO path provided.',
                console.error(err);
            });        
        },

        updating: function( status ) {
            return status == "starting" || status === "stopping";
        },
        delete_config(name ) {
            this.$buefy.dialog.confirm({
                title: 'Delete the image configuration',
                message: 'This will DELETE the ' + name + ' experiment. Are you sure you want to do this?',
                cancelText: 'Cancel',
                confirmText: 'Delete',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => {
                    this.isWaiting = true;
                    console.log("Delete image", name)
                    this.$http.delete(
                        'images/' + name
                    ).then(
                        response => {
                            if (response.status == 204) {
                                let img = this.images;
                                for (let i = 0; i < img.length; i ++){
                                    if (img[i].name == name) {
                                        img.splice(i, 1);
                                        break;
                                    }
                                }

                                this.images = [ ...img ];
                            }
                            this.isWaiting = false; 
                        }, err => {
                            this.errorNotification(err);
                            this.isWaiting = false;
                        }
                    )
                }
            })
        },
        //Build Modal / build functions 
        activateBuildModal(name){
            this.resetBuildModal()
            this.buildModal.name = name 
            this.buildModal.active = true 
        },
        resetBuildModal(){
            this.buildModal = {
                active: false,
                cache: false,
                dryrun: false,
                output: null,
                verbosity: 0,
                name: null,
            }
        },
        build(){
            this.isWaiting = true; 

            console.log(this.buildModal)

            let buildRequest = {
                name: this.buildModal.name,
                verbosity: this.buildModal.verbosity,
                cache: this.buildModal.cache,
                dryrun: this.buildModal.dryrun,
                output: this.buildModal.output,
            }
            this.$http.post('image/build', buildRequest, { timeout: 0 } 
                ).then(
                _ => {          
                    console.log("Build request sumbitted correctly")  
                    this.isWaiting = false;
                }, err => {
                    this.errorNotification(err);
                    this.isWaiting = false;
                }
                );

            this.buildModal.active = false;
            this.resetBuildModal()
            this.updateImages()
        },
        isEmpty(value){
            return (
                value == null ||
                (typeof value == 'string' && value.length === 0) ||
                value == []
            )
        }

    },
    data() {
        return {
            defaults: {},
            images: [],
            isWaiting: true,
            searchName: '',
            filtered: null,
            createModal: {
                active: false,
                name: null,
                type: null,
                nameErrType: null,
                nameErrMsg: null,
                osErrType: null,
                osErrMsg: null,
                tempPackage: null,

                winISOErrType: null,
                winISOErrError: null,

                enteredWinISO: null,
                winEditionOptions: null,

                winEditionErrType: null,
                winEditionErrError: null
            },
            buildModal: {
                active: false,
                cache: false,
                dryrun: false,
                output: null,
                verbosity: 0,
                name: null,

            },
            newImageForm: {
                name: null,
                deb_append: null,
                format: null,
                install_media: null,
                mirror: null,
                os: null,
                release: null,
                edition: null,
                size: null,
                variant: null,


                //list flags
                overlays: [],
                packages: [],
                script_order: [],
                script_paths: [],


                //boolean flags
                cache: false,
                compress: false,
                include_miniccc: false,
                include_protonuke: false,
                ramdisk: false,
                verbose_logs: false,
                

                //don't actually change
                scripts: {},
                
                
            }
        };
    }
}
</script>
<style scoped>
    div.autocomplete >>> a.dropdown-item {
        color: #383838 !important;
    }

    button.action {
        margin-right: 5px;
    }

    a.action {
        margin-right: 5px;
    }

    p.card-header-title {
        margin: 0;
    }

</style>