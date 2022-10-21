module bovine_vrese::bv_config {

    public fun admin_address(): address {
        @bovine_vrese
    }

    #[test]
    fun addresses() {
        admin_address();
    }

}