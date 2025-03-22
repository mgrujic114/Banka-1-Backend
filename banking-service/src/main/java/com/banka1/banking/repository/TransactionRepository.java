package com.banka1.banking.repository;

import com.banka1.banking.models.Account;
import com.banka1.banking.models.Transaction;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface TransactionRepository extends JpaRepository<Transaction, Long> {
    List<Transaction> findByFromAccountId(Account fromAccountId);
    List<Transaction> findByToAccountId(Account toAccountId);
    @Query("SELECT DISTINCT t FROM Transaction t WHERE t.fromAccountId IN :accounts OR t.toAccountId IN :accounts")
    List<Transaction> findByAccounts(@Param("accounts") List<Account> accounts);

}
