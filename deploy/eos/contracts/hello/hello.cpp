#include <eosio/eosio.hpp>
#include <string>

using namespace eosio;
using std::string;

class [[eosio::contract]] hello : public contract {
  public:
      using contract::contract;

      [[eosio::action]]
      void hi(string user) {
         print( "Hello, ", user);
         return;
      }
};