
<template>
    <div class="content">
        <b-modal :active.sync="createModal.active" :on-cancel="resetCreateModal" has-modal-card>
            <div class="modal-card" style="width:25em">
                <header class="modal-card-head">
                    <p class="modal-card-title">Create a New Image</p>
                </header>
                <section class="modal-card-body">
                    <b-field label="Image Name" 
                        :type="createModal.nameErrType" 
                        :message="createModal.nameErrMsg">
                        <b-input type="text" v-model="newImageForm.name"></b-input>
                    </b-field>
                    <b-field label="OS">
                        <b-input type="text" v-model="newImageForm.image.os"></b-input>
                    </b-field>
                    <b-field label="Variant">
                        <b-input type="text" v-model="newImageForm.image.variant"></b-input>
                    </b-field>
                    <b-field label="Release">
                        <b-input type="text" v-model="newImageForm.image.release"></b-input>
                    </b-field>
                    <b-field label="Format">
                        <b-input type="text" v-model="newImageForm.image.format"></b-input>
                    </b-field>
                    <b-field label="Size">
                        <b-input type="text" v-model="newImageForm.image.size"></b-input>
                    </b-field>
                    <b-field>
                        <b-checkbox v-model="newImageForm.image.compress" style="color:black">
                            Compress
                        </b-checkbox>
                    </b-field>
                    <b-field>
                        <b-checkbox v-model="newImageForm.image.ramdisk" style="color:black">
                            Ramdisk
                        </b-checkbox>
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
                            type="is-info">
                        </b-taginput>
                    </b-field>

                </section>

                <footer class="modal-card-foot buttons is-right">
                    <button class="button is-light" :disabled="!validate()" @click="create">Create Experiment</button>
                </footer>
            </div>
        </b-modal>


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
import ImageBuilderTagField from './ImageBuilderTagField.vue';

export default {
    beforeDestroy() {
        clearInterval(this.update);
    },
    created() {
        console.log("Created");
        this.isWaiting = false;
        // this.images = [];
        this.update_defaults();
    },
    methods: {
        get_image_list() {
            this.$http.get('image/list').then(response => {
                response.json().then(state => {
                    console.log(state);
                    this.images = state;
                });
            }, err => {
                console.log("Error found", err);
                this.errorNotification(err);
            });
        },

        update_defaults() {
            this.newImageForm.name = ""
            this.$http.get('image/create').then(response => {
                response.json().then(state => {
                    this.createModal.info = state;
                    this.newImageForm.image = state;
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
            return true;
        },
        resetCreateModal(){
            this.update_defaults();
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
