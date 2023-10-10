
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
                        <b-checkbox v-model="newImageForm.image.include_miniccc" style="color:black">
                            Include Miniccc
                        </b-checkbox>
                    </b-field>
                    <b-field label="Packages">
                        <b-taginput
                            v-model="newImageForm.image.packages"
                            placeholder="Add a package"
                            type="is-light">
                        </b-taginput>
                    </b-field>

                    <!-- <b-field label="Scripts">
                        <b-taginput
                            v-model="newImageForm.image.scripts"
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
                                        placeholder="Add a script">
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
        <hr>
        Welcome to the image builder page!
        <div id="test-div">
            <button @click="get_image_list()">Test GET /image/list</button>
            <br>
            <button @click="show_modal()">Show Modal</button>
        </div>
        <b-loading :is-full-page="true" :active.sync="isWaiting" :can-cancel="false"></b-loading>
    </div>
</template>

<script>

export default {
    beforeDestroy() {
        clearInterval(this.update);
    },
    created() {
        console.log("Created");
        this.isWaiting = false;
        // this.images = [];
        this.clear_form();
    },
    methods: {
        get_image_list() {
            this.$http.get(`image/list`).then(response => {
                response.json().then(state => {
                    console.log(state);
                    this.images = state;
                });
            }, err => {
                console.log("Error found", err);
                this.errorNotification(err);
            });
        },
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
            console.log("populating default values")
            this.$http.get(`image/create?os=${this.newImageForm.image.os}`).then(response => {
                response.json().then(state => {
                    console.log(this.newImageForm.image, state)
                    this.newImageForm.image = state
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
            }
            else {
                this.createModal.nameErrType = null;
                this.createModal.nameErrMsg = null;
            }

            if (this.newImageForm.image.os == null) {
                return false
            } 
            return true;
        },
        resetCreateModal(){
            this.clear_form();
        
        },
        create(){ 
            console.log(this.newImageForm)
            this.isWaiting = true;
            this.$http.post('image/create', this.newImageForm, { timeout: 0 } 
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
        }

    },
    computed: {},
    data() {
        return {
            defaults: {},
            images: [],
            isWaiting: true,
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
        };
    }
}
</script>
