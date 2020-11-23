<template>
<div>
    <!-- --- Top level search bar, search button, and reset button --- -->
    <div class="row justify-content-center">
        <div class="col-md-6 list">
            <form v-on:submit.prevent @submit="search">
                <div class="input-group mb-3">
                    <input ref="name" type="text" class="form-control fs-20" placeholder="Search by name" v-model="name" />
                    <div class="input-group-append">
                        <button class="btn btn-outline-primary ml-1 fs-20" type="submit">
                            Search
                        </button>
                        <button class="btn btn-outline-secondary fs-20" type="button" @click="reset">
                            Reset
                        </button>
                    </div>
                </div>
            </form>
        </div>
    </div>
    <!-- --- End top level search bar area ---  -->
</div>
</template>

<script>
export default {
    name: "search-bar",
    props: {
        method: {
            type: Function
        },
    },
    methods: {
        search() {
            if (this.name === "") {
                alert("Please enter a name to search for.");
                this.$refs.name.focus();
                return;
            }

            this.$emit('search', this.name);
            this.name = "";
        },
        reset(){
            this.name = "";
            this.$emit('reset');
        }
    },
    data() {
        return {
            name: ""
        };
    }
};
</script>
