`timescale 1ns / 1ps

module tb_top;

  //clock and reset signal declaration
  logic tb_clk, reset;
  logic [31:0] tb_WB_Data;

  // 100 MHz clock
  always #10 tb_clk = ~tb_clk;

  //reset Generation
  initial begin
    tb_clk = 0;
    reset = 1;
    #25 reset =0;
  end


  riscv riscV(
    .clk(tb_clk),
    .reset(reset),
    .WB_Data(tb_WB_Data)
  );

  task skip();
    @(negedge tb_clk);
    $display("SKIP [Cycle %2d - %3d ns] Got %x", ($time-1) / 20, $time, tb_WB_Data);
  endtask

  task test(int signed expected);
    @(negedge tb_clk);
    assert (tb_WB_Data==expected)
      $display("GOOD [Cycle %2d - %3d ns] Got %x", ($time-1) / 20, $time, tb_WB_Data);
    else
      $error("[Cycle %2d - %3d ns] Expected %x, but got %x", ($time-1) / 20, $time, expected, tb_WB_Data);
  endtask

  initial begin
    $system("echo; echo");
    @(posedge tb_clk);

    //
    // Cycle 0
    // Time 20ns
    //
    // FE - and x0, x0, x0
    skip();

    //
    // Cycle 1
    // Time 40ns
    //
    // FE - addi x1, x0, 1
    // DE - and x0, x0, x0
    skip();

    //
    // Cycle 2
    // Time 60ns
    //
    // FE - addi x2, x0, 1 -> 0x1
    // DE - addi x1, x0, 1
    // ME - and x0, x0, x0
    skip();

    //
    // Cycle 3
    // Time 80ns
    //
    // FE - jalr x3, 4(x0) -> 0x14
    // DE - addi x2, x0, 1 -> 0x1
    // ME - addi x1, x0, 1
    // EX - and x0, x0, x0
    skip();

    //
    // Cycle 4
    // Time 100ns
    //
    // DE - jalr x3, 4(x0) -> 0x14
    // ME - addi x2, x0, 1 -> 0x1
    // EX - addi x1, x0, 1
    // WB - and x0, x0, x0
    skip();

    //
    // Cycle 5
    // Time 120ns
    //
    // ME - jalr x3, 4(x0) -> 0x14
    // EX - addi x2, x0, 1 -> 0x1
    // WB - addi x1, x0, 1
    skip();

    //
    // Cycle 6
    // Time 140ns
    //
    // EX - jalr x3, 4(x0) -> 0x14
    // WB - addi x2, x0, 1 -> 0x1
    test(32'h1);

    //
    // Cycle 7
    // Time 160ns
    //
    // WB - jalr x3, 4(x0) -> 0x14
    test(32'h14);



    $system("echo; echo; echo");
    $finish;
  end

endmodule
