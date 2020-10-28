<template>
<div>
    <nav aria-label="Pagination">
        <ul class="pagination mt-2 justify-content-center">
            <li v-if="loaded" :key="'prev'" :class="{'disabled':currentPage === 1 || totalPages === 0}" class="page-item previous-item">
                <button class="page-link" @click="currentPage--">Prev</button>
            </li>
            <li v-for="page in pageArray" :key="page" class="page-item" :class="{'active':(currentPage === page)}">
                <button class="page-link" @click="currentPage = page">{{page}}</button>
            </li>
            <li v-if="loaded" :key="'next'" :class="{'disabled':currentPage === totalPages || totalPages === 0}" class="page-item next-item">
                <button class="page-link" @click="currentPage++">Next</button>
            </li>
        </ul>
    </nav>
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
        reset() {
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
