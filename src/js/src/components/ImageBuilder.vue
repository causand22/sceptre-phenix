
<template>
    <div class="content">
        <b-modal :active.sync="createModal.active" :on-cancel="resetCreateModal" has-modal-card>
            <div class="modal-card" style="width:40em">
                <header class="modal-card-head">
                    <p class="modal-card-title">Create a New Image</p>
                </header>
                <section class="modal-card-body">

                    <b-field label="Image Name" 
                        :type="createModal.nameErrType" 
                        :message="createModal.nameErrMsg">
                        <b-input type="text" v-model="newImageForm.name"></b-input>
                    </b-field>


                    <b-field label="OS"
                        :type="createModal.osErrType"
                        :message="createModal.osErrMsg">
                        <div class="level">
                            <div class="level-left">
                                <b-select placeholder="Select an OS" v-model="newImageForm.image.os">
                                    <option value="linux">Linux</option>
                                    <option value="windows">Windows</option>
                                </b-select>
                            </div>
                            <div class="level-right">
                                <button class="button is-light" @click="populate_defaults()">Populate Defaults</button>
                            </div>
                        </div>
                    </b-field>

                    <b-collapse class="card" animation="slide" :open="false" v-if="newImageForm.image.os == 'linux'">
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
                                    <b-input type="text" v-model="newImageForm.image.variant"></b-input>
                                </b-field>
                                <b-field label="Release">
                                    <b-input type="text" v-model="newImageForm.image.release"></b-input>
                                </b-field>
                                <b-field label="Mirror">
                                    <b-input type="text" v-model="newImageForm.image.mirror"></b-input>
                                </b-field>
                                <b-field label="Install Media">
                                    <b-input type="text" v-model="newImageForm.image.install_media"></b-input>
                                </b-field>
                                <b-field label="Packages">
                                    <b-taginput
                                        v-model="newImageForm.image.packages"
                                        placeholder="Add a package"
                                        type="is-light">
                                    </b-taginput>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>
                    <b-collapse class="card" animation="slide" :open="false" v-if="newImageForm.image.os == 'windows'">
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
                                <b-field label="Install Media ISO">
                                    <b-input type="text" v-model="newImageForm.image.install_media"></b-input>
                                </b-field>
                                <b-field label="Edition">
                                    <b-input type="text" v-model="newImageForm.image.release"></b-input>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>


                    <b-field label="Format">
                        <b-input type="text" v-model="newImageForm.image.format"></b-input>
                    </b-field>
                    <b-field label="Size">
                        <b-input type="text" v-model="newImageForm.image.size"></b-input>
                    </b-field>


                    <b-field>
                        <b-checkbox v-model="newImageForm.image.include_protonuke" style="color:black">
                            Include Protonuke
                        </b-checkbox>
                    </b-field>
                    <b-field>
                        <b-checkbox v-model="newImageForm.image.include_miniccc"
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
                                    <b-input type="textarea" v-model="newImageForm.image.deb_append"></b-input>
                                    <!-- <b-tooltip label="Additional arguments to debootstrap"
                                        type="is-light"
                                        multilined>
                                        <b-icon icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip> -->
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.image.verbose_logs" style="color:black">Verbose Logs</b-checkbox>
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.image.compress" style="color:black">Compress</b-checkbox>
                                    <b-tooltip label="Compress image after creation (does not apply to raw image)" 
                                                    type="is-light" 
                                                    multilined>
                                        <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip>
                                </b-field>
                                <b-field>
                                    <b-checkbox v-model="newImageForm.image.ramdisk" style="color:black">Ram / Disk</b-checkbox>
                                    <b-tooltip label="Create a kernel/initrd pair in addition to a disk image" 
                                                    type="is-light" 
                                                    multilined>
                                        <b-icon  icon="question-circle" style="color:#383838"></b-icon>
                                    </b-tooltip>                                
                                </b-field>
                                <b-field label="Scripts">
                                    <b-taginput
                                        v-model="newImageForm.image.script_paths"
                                        placeholder="Add a script (include full path)">
                                    </b-taginput>
                                </b-field>
                                <b-field label="Overlays">
                                    <b-taginput
                                        v-model="newImageForm.image.overlays"
                                        placeholder="Add an overlay (include full path)">
                                    </b-taginput>
                                </b-field>
                            </div>
                        </div>
                    </b-collapse>
                </section>

                <footer class="modal-card-foot">
                    <button class="button is-danger mr-auto" @click="clear_form">Reset Form</button>
                    <button class="button is-light" :disabled="!validate()" @click="create">Create Image</button>
                </footer>
            </div>
        </b-modal>

        <template v-if="images.length == 0">
        <section class="hero is-light is-bold is-large">
            <div class="hero-body">
            <div class="container" style="text-align: center">
                <!-- <h1 class="title">
                There are no experiments!
                </h1> -->
                <!-- #TODO: add role allowed-->
                <b-button type="is-success" outlined @click="createModal.active = true">
                    Create An Image
                </b-button>
            </div>
            </div>
        </section>
        </template>
        <template v-else>
            <hr>
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
                <b-tooltip label="create a new image config" type="is-light" multilined>
                    <button class="button is-light" @click="updateImages(); createModal.active = true">
                    <b-icon icon="plus"></b-icon>
                    </button>
                </b-tooltip>
            </b-field>
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
                    <b-table-column field="updated" label="Last Updated" width="200" sortable v-slot="props">
                        {{ props.row.updated }}
                    </b-table-column>
                    <b-table-column field="size" label="Size" width="200" sortable v-slot="props">
                        {{ props.row.size }}
                    </b-table-column>
                    <b-table-column field="os" label="Os" width="200" sortable v-slot="props">
                        {{ props.row.os }}
                    </b-table-column>
                    <b-table-column field="format" label="Format" width="200" sortable v-slot="props">
                        {{ props.row.format }}
                    </b-table-column>
                    <b-table-column label="Actions" width="150" centered v-slot="props">
                        <b-tooltip label="Build" type="is-light">
                            <button
                                class="button is-light is-small action"
                                :disabled="updating( props.row.status )"
                                @click="build(props.row.name)">
                                Build
                            </button>
                        </b-tooltip>
                        <b-tooltip label="Edit" type="is-light">
                            <button
                                class="button is-light is-small action"
                                :disabled="updating( props.row.status )"
                                @click="edit(props.row.name)">
                            <b-icon icon="edit"></b-icon>
                            </button>
                        </b-tooltip>
                        <b-tooltip label="Delete" type="is-light">
                            <button 
                                class="button is-light is-small action" 
                                :disabled="updating( props.row.status )" 
                                @click="del(props.row.name)">
                            <b-icon icon="trash"></b-icon>
                            </button>
                        </b-tooltip>



                    </b-table-column>
                </b-table>
            </div>
        </template>
        <button @click="updateImages">Click to update images</button>
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
        this.clear_form();
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
        }
    },
    methods: {
        clear_form(){
            this.newImageForm = {
                name: null,
                image: {
                    cache: false,
                    compress: false,
                    deb_append: null,
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
            }
        },
        populate_defaults() {
            // console.log("populating default values")
            this.$http.get(`image/create?os=${this.newImageForm.image.os}`).then(response => {
                response.json().then(state => {
                    // console.log("Log 1", this.newImageForm, state)

                    function isEmpty(value){
                        return (
                            value == null ||
                            (typeof value == 'string' && value.length === 0) ||
                            value == []
                        )
                    }
                    for (const [key, value] of Object.entries(state)) {
                        if (!isEmpty(value)) {
                            // console.log(key)
                            this.newImageForm.image[key] = value
                            // console.log(key, this.newImageForm.image[key], value)
                        }
                    }
                    // console.log(this.newImageForm, state)

                    // this.newImageForm = state
                });
            }, err => {
                console.error(err);
            });
        },

        show_modal() {
            this.createModal.active = true;
        },
        validate() {
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

            if (this.newImageForm.image.os == null) {
                return false
            } 

            // const sizeRegex = new RegExp('[0-9]+[M|G]')
            // if (! sizeRegex.test(this.newImageForm.size)){
            //     return false 
            // }
            return true;
        },
        resetCreateModal(){
            this.clear_form();
        
        },
        create(){ 
            // console.log(this.newImageForm)
            this.isWaiting = true;

            let form = this.newImageForm.image
            form.name = this.newImageForm.name

            this.$http.post('image/create', form, { timeout: 0 } 
                ).then(
                _ => {          
                    console.log("Request sumbitted correctly")  
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
        updateImages(){
            this.images = []
            function convertToImage(image) {

                var imageObj = {
                    name: image.Metadata.name,
                    size: image.Spec.size,
                    os: image.Spec.os,
                    format: image.Spec.format,
                    status: "stopped",
                    updated: image.Metadata.updated
                }
                return imageObj
            }

            this.$http.get('images').then(
                response => {
                    response.json().then( state => {
                        // console.log(state)
                        for (const [key, value] of Object.entries(state)) {
                            // console.log(key, value)
                            this.images.push(convertToImage(value))
                        }

                    })
                }
            )
        },
        updating: function( status ) {
            return status == "starting" || status === "stopping";
        },
        del (name ) {
            this.$buefy.dialog.confirm({
                title: 'Delete the image configuration',
                message: 'This will DELETE the ' + name + 'experiment. Are you sure you want to do this?',
                cancelText: 'Cancel',
                confirmText: 'Delete',
                type: 'is-danger',
                hasIcon: true,
                onConfirm: () => {
                    this.isWaiting = true;
                    console.log("Delete image ")
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
        build(name){
            console.log("Building image:", name)
            console.error("Implement function build")
        },
        edit(name){
            console.log("Editing image:", name)
            console.error("Implement function edit")
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
                nameErrType: null,
                nameErrMsg: null,
                osErrType: null,
                osErrMsg: null,
                tempPackage: null,
            },
            newImageForm: {
                name: null,
                image: {
                    deb_append: null,
                    format: null,
                    install_media: null,
                    mirror: null,
                    os: null,
                    release: null,
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